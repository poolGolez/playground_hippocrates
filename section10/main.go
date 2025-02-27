package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Creating Gin Server for Loans...")
	server := gin.Default()
	server.SetTrustedProxies(nil)

	server.GET("/", helloWorld)

	server.Run()
}

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}
