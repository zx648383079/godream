package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	connection *gorm.DB
)

func init() {
	connection = connect()
}

func GetDB() *gorm.DB {
	return connection
}

func connect() *gorm.DB {
	url := os.Getenv("DATABASE")

	conn, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	return &conn
}