package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB = New()
)

/**
 * 设置数据库连接
 * @param diver string
 */
func New() *gorm.DB {
	url := os.Getenv("DATABASE")

	DB, err := gorm.Open("mysql", url)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}

	return DB
}
