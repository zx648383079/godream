package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS CORS Handler
var CORS gin.HandlerFunc

func initCORS() {
	CORS = cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"x-xq5-jwt", "Content-Type", "Origin", "Content-Length"},
		ExposeHeaders:    []string{"x-xq5-jwt"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
