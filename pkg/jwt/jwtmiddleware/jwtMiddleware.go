package jwtmiddleware

import (
	"fmt"
	"net/http"
	"url-shortener/pkg/jwt/jwtservice"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwtservice.JWTAuthService().ValidateToken(authHeader)
	if err != nil && !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println(err)
	} else {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
	}
}
