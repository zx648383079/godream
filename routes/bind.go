package routes

import (
	"github.com/gin-gonic/gin"
)

type (
	GroupFunc    = func(app *gin.RouterGroup)
	GroupFuncMap = map[string]GroupFunc
)

func RegisterMap(app *gin.Engine, items GroupFuncMap) {
	for path, v := range items {
		g := app.Group(path)
		v(g)
	}
}
