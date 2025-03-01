package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	server.GET("/loans/:id", fetchLoan)
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

func fetchLoan(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	loan := loans.FetchById(id)
	if loan == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, *loan)
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
