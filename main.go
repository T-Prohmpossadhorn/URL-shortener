package main

import (
	"fmt"

	handlers "url-shortener/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handlers.DefaultHandler)

	err := r.Run(":5000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
