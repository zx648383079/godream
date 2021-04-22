package database

import (
	"fmt"

	"zodream.cn/godream/configs"
	"zodream.cn/godream/utils/search"

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

func Where(query interface{}, args ...interface{}) *gorm.DB {
	return DB.Where(query, args...)
}

func Table(table string) *gorm.DB {
	return DB.Table(table)
}

func Model(value interface{}) *gorm.DB {
	return DB.Model(value)
}

func Raw(sql string, values ...interface{}) *gorm.DB {
	return DB.Raw(sql, values...)
}

func Search(columns []string, value string) *gorm.DB {
	return search.Where(DB, columns, value)
}

func First(dest interface{}, conds ...interface{}) *gorm.DB {
	return DB.First(dest, conds...)
}

func Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return DB.Find(dest, conds...)
}

func Create(value interface{}) *gorm.DB {
	return DB.Create(value)
}

func Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return DB.Delete(value, conds...)
}
