package routes

import (
	"net/http"

	"example.com/gin/loaney/loans"
	"github.com/gin-gonic/gin"
)

func listLoans(ctx *gin.Context) {
	loans := loans.FetchAll()
	ctx.JSON(http.StatusOK, loans)
}

func fetchLoan(ctx *gin.Context) {
	loan, _ := ctx.Get("X-Loan")
	ctx.JSON(http.StatusOK, loan)
}

func createLoan(ctx *gin.Context) {
	var loan loans.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Malformed request", "message": err})
		return
	}

	loans.Save(&loan)

	ctx.JSON(http.StatusCreated, loan)
}

func updateLoan(ctx *gin.Context) {
	loan, _ := ctx.Get("X-Loan")
	ctx.JSON(http.StatusOK, loan)
}
