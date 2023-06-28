package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/go-bank/model/dto"
	"github.com/rafialariq/go-bank/service"
	"github.com/rafialariq/go-bank/utility"
)

type LoginController struct {
	loginService service.LoginService
}

// func NewLoginController(r *gin.RouterGroup, ls service.LoginService) *LoginController {
// 	controller := LoginController{
// 		loginService: ls,
// 	}
// 	r.GET("/login", controller.LoginHandler)
// 	return &controller
// }

func NewLoginController(ls service.LoginService) *LoginController {
	controller := LoginController{
		loginService: ls,
	}

	return &controller
}

func (l *LoginController) LoginHandler(ctx *gin.Context) {
	var user dto.LoginDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	res, err := l.loginService.FindUser(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	// set cookie expiration time
	authDuration, _ := strconv.Atoi(utility.DotEnv("AUTH_DURATION", ".env"))
	expirationTime := time.Now().Add(time.Minute * time.Duration(authDuration)).Unix()

	// set JWT token to cookie
	ctx.SetCookie("token", res, int(expirationTime), "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"token": res,
	})
}
