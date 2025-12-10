package route

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"ByTeora-Pos-Backend-App/controller"
)

func RegisterRoutes(r *gin.Engine) {

	// Route default (homepage API)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Byteora API",
		})
	})

	// API Group
	api := r.Group("/api/v1")

	// User Routes
	api.POST("/users", controller.CreateUser)
	api.POST("/auth/login", controller.AuthLogin)
}