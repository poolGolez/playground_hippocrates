package routes

import (
	"example.com/gin/loaney/users"
	"github.com/gin-gonic/gin"
)

func GetCurrentUser(ctx *gin.Context) users.CurrentUser {
	data, exists := ctx.Get("X-CurrentUser")
	if !exists {
		panic("Current user wasn't injected within context")
	}

	currentUser, ok := data.(users.CurrentUser)

	if !ok {
		panic("Corrupted current user")
	}

	return currentUser
}
