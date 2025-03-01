package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/loans", listLoans)
	server.GET("/loans/:id", fetchLoan)
	server.POST("/loans", createLoan)
}
