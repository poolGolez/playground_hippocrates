package routes

import (
	"net/http"
	"strconv"

	"example.com/gin/loaney/loans"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/loans", listLoans)
	server.POST("/loans", createLoan)

	group := server.Group("/loans/:id").Use(getLoan)
	{
		group.GET("", fetchLoan)
		group.PUT("", updateLoan)
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
