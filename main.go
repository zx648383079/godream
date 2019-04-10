package main

import (
	"os"

	"zodream/router"

	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	addr := os.Getenv("PORT")
	router.Register(app)
	app.Run(iris.Addr(addr))
}
