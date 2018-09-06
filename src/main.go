package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/gzip"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"./database"
	"flag"
	"./chat"
)

const defaultPort = "8080"

func main() {
	db := database.GetDB()
	defer db.Close()
	flag.Parse()
	hub := newHub()
	chat.HubBox = hub
	go hub.run()
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.GET("/ws", chat.Chat)
	
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	router.Run(port)
}