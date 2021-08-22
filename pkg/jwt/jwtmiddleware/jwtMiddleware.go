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
	tokenString := authHeader[:]
	token, err := jwtservice.JWTAuthService().ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
	} else {
		fmt.Println("testing")
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
