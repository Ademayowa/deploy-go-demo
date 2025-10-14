package routes

import (
	"net/http"

	"github.com/Ademayowa/deploy-go-demo/models"
	"github.com/gin-gonic/gin"
)

// Create a property
func createProperty(context *gin.Context) {
	var property models.Property

	err := context.ShouldBindJSON(&property)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse property data"})
		return
	}
	property.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "property created", "property": property})
}

// Fetch all properties
func getProperties(context *gin.Context) {

	// Get all properties
	properties, err := models.GetAllProperties()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch properties"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"properties": properties})
}

// Fetch single property
func getProperty(context *gin.Context) {
	propertyID := context.Param("id")

	property, err := models.GetPropertyByID(propertyID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch property"})
		return
	}

	context.JSON(http.StatusOK, property)
}

// Delete property
func deleteProperty(context *gin.Context) {
	propertyID := context.Param("id")

	property, err := models.GetPropertyByID(propertyID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch property"})
		return
	}

	err = property.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete property"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "property deleted"})
}

// Update property
func updateProperty(context *gin.Context) {
	// Extract property ID from the URL
	propertyID := context.Param("id")

	// Parse the request body to get the updated property data
	var updatedProperty models.Property
	if err := context.ShouldBindJSON(&updatedProperty); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Update property in the database
	err := models.UpdateProperty(propertyID, updatedProperty)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not update property"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "property updated"})
}
