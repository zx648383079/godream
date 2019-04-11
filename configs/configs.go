package configs

import "github.com/jinzhu/configor"

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

func Init(file string) {
	configor.Load(&Config, file)
}
