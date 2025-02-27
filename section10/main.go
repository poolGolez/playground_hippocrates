package main

import (
	"fmt"
	"net/http"

	"example.com/gin/loaney/loan"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Creating Gin Server for Loans...")
	server := gin.Default()
	server.SetTrustedProxies(nil)

	server.GET("/", helloWorld)
	server.GET("/loans", listLoans)

	server.Run()
}

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func listLoans(ctx *gin.Context) {
	loans := loan.FetchAllLoans()
	ctx.JSON(http.StatusOK, loans)
}
