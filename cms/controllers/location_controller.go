package controllers

import (
	"net/http"

	"cqrs_demo/database"
	"cqrs_demo/models"

	"github.com/gin-gonic/gin"
)

// Get all locations
func GetLocations(c *gin.Context) {
	var locations []models.Location
	result := database.GetDB().Find(&locations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": locations})
}

// Get location by ID
func GetLocation(c *gin.Context) {
	id := c.Param("id")
	var location models.Location

	result := database.GetDB().First(&location, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": location})
}

// Create location
func CreateLocation(c *gin.Context) {
	var location models.Location

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Create(&location)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": location})
}

// Update location
func UpdateLocation(c *gin.Context) {
	id := c.Param("id")
	var location models.Location

	// Check if location exists
	if err := database.GetDB().First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	var updateData models.Location
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Model(&location).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": location})
}

// Delete location
func DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	var location models.Location

	result := database.GetDB().Delete(&location, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
