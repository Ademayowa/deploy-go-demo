package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/properties", createProperty)
	router.GET("/properties", getProperties)
	router.GET("/properties/:id", getProperty)
	router.DELETE("/properties/:id", deleteProperty)
	router.PUT("/properties/:id", updateProperty)
}
