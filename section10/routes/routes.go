package routes

import (
	"net/http"
	"strconv"

	"example.com/gin/loaney/loans"
	mw "example.com/gin/loaney/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/users/register", RegisterUser)
	server.POST("/users/login", Login)

	authGroup := server.Group("/loans")
	authGroup.Use(mw.Authenticate)
	authGroup.GET("/", listLoans)
	authGroup.POST("/", createLoan)

	group := authGroup.Group("/:id").Use(getLoan)
	{
		group.GET("", fetchLoan)
		group.PUT("", updateLoan)
		group.DELETE("", deleteLoan)
	}
}

func getLoan(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid loan ID"})
		return
	}

	loan := loans.FetchById(id)
	if loan == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}
	ctx.Set("X-Loan", loan)
}
