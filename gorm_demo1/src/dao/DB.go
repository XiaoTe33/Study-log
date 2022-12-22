package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_demo1/src/model"
)

var DB = DBInit()

func DBInit() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/im_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("fail to connect with mysql")
		return nil
	}
	err = db.AutoMigrate(&model.Student{})
	if err != nil {
		fmt.Printf("err = %v", err)
		return nil
	}
	return db
}
