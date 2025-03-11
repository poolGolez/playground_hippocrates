package middlewares

import (
	"net/http"
	"strings"

	"example.com/gin/loaney/users"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	authBearer := ctx.GetHeader("Authorization")
	tokens := strings.Fields(authBearer)

	if len(tokens) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	jwt := tokens[len(tokens)-1]

	currentUser, err := users.DecodeJwt(jwt)

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Set("X-CurrentUser", *currentUser)

	ctx.Next()
}
