package routes

import (
	"net/http"

	"example.com/gin/loaney/users"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var params users.RegisterUserParams
	ctx.ShouldBindJSON(&params)

	user, _ := users.Register(params)

	ctx.JSON(http.StatusOK, gin.H{"Hello": "world", "user": user})
}
