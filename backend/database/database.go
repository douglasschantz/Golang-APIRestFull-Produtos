package database

import (
	"github.com/schantz/web/go-api-produtos/backend/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(localhost:3306)/northwind?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	utils.ErrorPanic(err)

	db = conn

	//return db
}

func ReturnDB() *gorm.DB {
	return db
}
