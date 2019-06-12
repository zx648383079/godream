package database

import (
	"fmt"

	"zodream/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
* 数据库
 */
var DB *gorm.DB

/*
 * 设置数据库连接
 */
func New() *gorm.DB {
	url := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configs.Config.Db.User, configs.Config.Db.Password, configs.Config.Db.Schema)

	db, err := gorm.Open(configs.Config.Db.Driver, url)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
		return nil
	}
	DB = db
	return db
}
