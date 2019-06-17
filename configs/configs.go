package configs

import (
	"fmt"

	"github.com/jinzhu/configor"
)

// Config 系统配置信息
var Config = struct {
	Host    string `default:"localhost"`
	Port    string `default:"8080"`
	Debug   bool
	Favicon string `default:"./assets/favicon.ico"`
	Asset   string `default:"./assets"`
	View    string `default:"./templates"`
	Db      struct {
		Driver   string `default:"mysql"`
		Host     string `default:"localhost"`
		Port     string `default:"3306"`
		User     string `default:"root"`
		Password string `default:""`
		Schema   string
	}
}{}

// Init 配置初始化
func Init(file string) {
	configor.Load(&Config, file)
}

// Host 获取host
func Host() string {
	return fmt.Sprintf("%s:%s", Config.Host, Config.Port)
}
