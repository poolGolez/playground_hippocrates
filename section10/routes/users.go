package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Hello": "world"})
}
