package main

import (
	"url-shortener/config"
	"url-shortener/internal/handlers"
	"url-shortener/internal/store/redis"
	"url-shortener/pkg/jwt/jwthandlers"
	"url-shortener/pkg/jwt/jwtmiddleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.FromFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	redis, err := redis.New(*config)
	if err != nil {
		panic("Failed to start storage connection " + err.Error())
	}
	defer redis.Close()

	handlers.Setupstore(&redis, *config)

	jwthandlers.Initialize(*config)

	r := gin.Default()
	r.GET("/", handlers.DefaultHandler)
	r.GET("/:shortUrl", handlers.HandleShortUrlRedirect)

	r.POST("/login", jwthandlers.GenerateToken)

	admin := r.Group("/admin")
	admin.Use(jwtmiddleware.AuthorizeJWT)
	{
		admin.GET("/test", jwthandlers.TestJwt)
		admin.POST("/", handlers.CreateShortUrl)
		admin.GET("/:shortUrl", handlers.GetShortUrlInfo)
		admin.DELETE("/:shortUrl", handlers.DeleteShortUrl)
	}

	err = r.Run(":" + config.Server.Port)
	if err != nil {
		panic("Failed to start the web server - Error: " + err.Error())
	}
}
