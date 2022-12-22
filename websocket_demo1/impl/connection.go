package impl

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsCoon    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsCoon:    wsConn,
		inChan:    make(chan []byte, 1024),
		outChan:   make(chan []byte, 1024),
		closeChan: make(chan byte, 1),
		mutex:     sync.Mutex{},
		isClosed:  false,
	}
	go conn.readLoop()
	go conn.writeLoop()
	return
}

func (conn *Connection) Close() {
	_ = conn.wsCoon.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
}

func (conn *Connection) readLoop() {
	for {
		_, data, err := conn.wsCoon.ReadMessage()
		if err != nil {
			fmt.Println("err!")
			conn.Close()
			return
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			conn.Close()
			return
		}
	}
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			conn.Close()
			return
		}
		if err = conn.wsCoon.WriteMessage(websocket.TextMessage, data); err != nil {
			fmt.Println("err")
			conn.Close()
		}
	}
}

func (conn *Connection) ReadMessage() ([]byte, error) {
	var data []byte
	var err error
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("conn closed")
	}
	return data, err
}

func (conn *Connection) WriteMessage(data []byte) error {
	var err error
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("conn closed")
	}
	return err
}
