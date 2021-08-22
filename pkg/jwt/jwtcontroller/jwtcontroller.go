package jwtcontroller

import (
	"url-shortener/pkg/jwt/jwtdto"
	"url-shortener/pkg/jwt/jwtservice"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService jwtservice.LoginService
	jWtService   jwtservice.JWTService
}

func LoginHandler(loginService jwtservice.LoginService,
	jWtService jwtservice.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential jwtdto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}
