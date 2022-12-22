package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
	"websocket_demo2/impl"
)

func main() {

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		upGrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("upGrade err!")
			return
		}
		conn, err := impl.InitConnection(wsConn)
		if err != nil {
			fmt.Println("InitConnection err!")
		}

		go func() {
			begin := time.Now().Unix()
			for {
				now := time.Now().Unix()
				t := strconv.Itoa(int(now - begin))
				err2 := conn.WriteMessage([]byte("webSocket已经连接" + t + "秒"))
				if err2 != nil {
					conn.Close()
				}
				time.Sleep(time.Second * 30)
			}
		}()

		var data []byte
		for {
			if data, err = conn.ReadMessage(); err != nil {
				conn.Close()
			}
			if err = conn.WriteMessage(data); err != nil {
				conn.Close()
			}
		}
	})
	_ = r.Run(":8081")
}
