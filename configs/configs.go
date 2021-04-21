package configs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/configor"
	"zodream.cn/godream/utils"
)

type SystemConfig struct {
	Host    string `default:"localhost"`
	Port    string `default:"8080"`
	Debug   bool
	Favicon string `default:"./assets/favicon.ico"`
	Asset   string `default:"./assets"`
	Upload  string `default:"upload"`
	View    string `default:"./templates"`
	Db      struct {
		Driver   string `default:"mysql"`
		Host     string `default:"localhost"`
		Port     string `default:"3306"`
		User     string `default:"root"`
		Password string `default:""`
		Schema   string
	}
	Auth struct {
		Key string `default:"zodream cn"`
	}
	Route struct {
		Deeplink string `default:"zodream"`
	}
}

// Config 系统配置信息
var Config = SystemConfig{}

// Init 配置初始化
func Init(file string) {
	configor.Load(&Config, file)
}

// Host 获取host
func Host() string {
	return fmt.Sprintf("%s:%s", Config.Host, Config.Port)
}

// UploadRandomFile 根据带点的拓展名生成随机路径
func UploadRandomFile(ext string) (string, string) {
	folder := "file"
	now := time.Now()
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".bmp", ".webp":
		folder = "image"
	case ".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
		".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid":
		folder = "video"
	}
	return UploadFile(folder + "/" + strconv.FormatInt(now.UnixNano(), 10) + ext)
}

// UploadRandomFileName 根据文件名生成随机的文件路径
func UploadRandomFileName(name string) (string, string) {
	return UploadRandomFile(utils.FileExtension(name))
}

// UploadFile 获取上传保存路径及网址
func UploadFile(name string) (string, string) {
	path := "/" + strings.Trim(Config.Upload, "/") + "/" + name
	return strings.TrimRight(Config.Asset, "/") + path, "/assets" + path
}
