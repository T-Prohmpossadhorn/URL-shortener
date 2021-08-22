package jwthandlers

import (
	"net/http"
	"url-shortener/pkg/jwt/jwtcontroller"
	"url-shortener/pkg/jwt/jwtservice"

	"github.com/gin-gonic/gin"
)

var loginService jwtservice.LoginService
var jwtService jwtservice.JWTService
var loginController jwtcontroller.LoginController

func Initialize() {
	loginService = jwtservice.StaticLoginService()
	jwtService = jwtservice.JWTAuthService()
	loginController = jwtcontroller.LoginHandler(loginService, jwtService)
}

func GenerateToken(c *gin.Context) {
	token := loginController.Login(c)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

func TestJwt(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
