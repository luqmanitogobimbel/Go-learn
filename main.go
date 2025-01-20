package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luqmanito/go-learn/controllers/userController"
	"github.com/luqmanito/go-learn/models"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/api/users", userController.Index)
	r.GET("/api/user/:id", userController.Show)
	r.POST("/api/user", userController.Create)
	r.PUT("/api/user/:id", userController.Update)
	r.DELETE("/api/user", userController.Delete)

	r.Run(":8081")
}
