package main

import (
	"os"
	"zodream/controllers"
	"zodream/modules/auth"
	"zodream/modules/gzo"
)

func (app *app) Register() {
	app.Get("/", controllers.Index)
	app.Get("/home", controllers.Index)
	app.PartyFunc("/auth", auth.Register)
	if os.Getenv("DEBUG") == "true" {
		app.PartyFunc("/gzo", gzo.Register)
	}
}
