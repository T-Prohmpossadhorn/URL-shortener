package jwtmiddleware

import (
	"fmt"
	"net/http"
	"url-shortener/pkg/jwt/jwtservice"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		token, err := jwtservice.JWTAuthService().ValidateToken(authHeader)
		if err != nil && !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println(err)
		} else {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
