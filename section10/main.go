package main

import (
	"fmt"
	"net/http"

	"example.com/gin/loaney/loans"
	"github.com/gin-gonic/gin"
)

func main() {
	loans.InitDB()

	fmt.Println("Creating Gin Server for Loans...")
	server := gin.Default()
	server.SetTrustedProxies(nil)

	server.GET("/", helloWorld)
	server.GET("/loans", listLoans)
	server.POST("/loans", createLoan)

	server.Run()
}

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func listLoans(ctx *gin.Context) {
	loans := loans.FetchAll()
	ctx.JSON(http.StatusOK, loans)
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
