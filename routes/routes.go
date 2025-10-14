package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/properties", CreateProperty)
	router.GET("/properties", GetProperties)
	router.GET("/properties/:id", GetProperty)
	router.DELETE("/properties/:id", DeleteProperty)
	router.PUT("/properties/:id", UpdateProperty)
}
