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
	config, err := config.FromFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	jwthandlers.Initialize(*config)

	r := gin.Default()
	r.GET("/", handlers.DefaultHandler)

	r.POST("/login", jwthandlers.GenerateToken)

	admin := r.Group("/admin")
	admin.Use(jwtmiddleware.AuthorizeJWT)
	{
		admin.GET("/test", jwthandlers.TestJwt)
	}

	err = r.Run(":" + config.Server.Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
