package main

import (
	"fmt"

	"url-shortener/config"
	"url-shortener/internal/handlers"
	"url-shortener/pkg/jwt/jwthandlers"
	"url-shortener/pkg/jwt/jwtmiddleware"

	"github.com/gin-gonic/gin"
)

func main() {
	jwthandlers.Initialize()

	r := gin.Default()
	r.GET("/", handlers.DefaultHandler)

	r.POST("/login", jwthandlers.GenerateToken)

	admin := r.Group("/admin")
	admin.Use(jwtmiddleware.AuthorizeJWT)
	{
		admin.GET("/test", jwthandlers.TestJwt)
	}

	err := r.Run(":" + config.PORT)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
