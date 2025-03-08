package routes

import (
	"net/http"

	"example.com/gin/loaney/users"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var params users.LoginParams
	ctx.ShouldBindJSON(&params)
	jwt := users.Login(params)

	ctx.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func RegisterUser(ctx *gin.Context) {
	var params users.RegisterUserParams
	ctx.ShouldBindJSON(&params)

	user, _ := users.Register(params)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
