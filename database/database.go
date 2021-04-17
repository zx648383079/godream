package database

import (
	"fmt"

	"zodream.cn/godream/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
* 数据库
 */
var DB *gorm.DB

/*
 * 设置数据库连接
 */
func InitDb() error {
	url := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configs.Config.Db.User, configs.Config.Db.Password, configs.Config.Db.Schema)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
		return err
	}
	DB = db
	return nil
}
