package main

import (
	"fmt"
	"net/http"

	"example.com/gin/loaney/db"
	"example.com/gin/loaney/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("Creating Gin Server for Loans...")
	server := gin.Default()
	server.SetTrustedProxies(nil)

	server.GET("/", helloWorld)
	routes.RegisterRoutes(server)
	server.Run()
}

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}
