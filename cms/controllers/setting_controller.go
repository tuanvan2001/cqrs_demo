package controllers

import (
	"net/http"

	"cqrs_demo/database"
	"cqrs_demo/models"

	"github.com/gin-gonic/gin"
)

// Get all settings
func GetSettings(c *gin.Context) {
	var settings []models.Setting
	result := database.GetDB().Preload("Location").Find(&settings)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": settings})
}

// Get setting by ID
func GetSetting(c *gin.Context) {
	id := c.Param("id")
	var setting models.Setting

	result := database.GetDB().Preload("Location").First(&setting, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": setting})
}

// Create setting
func CreateSetting(c *gin.Context) {
	var setting models.Setting

	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Create(&setting)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": setting})
}

// Update setting
func UpdateSetting(c *gin.Context) {
	id := c.Param("id")
	var setting models.Setting

	if err := database.GetDB().First(&setting, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		return
	}

	var updateData models.Setting
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Model(&setting).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": setting})
}

// Delete setting
func DeleteSetting(c *gin.Context) {
	id := c.Param("id")
	var setting models.Setting

	result := database.GetDB().Delete(&setting, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Setting deleted successfully"})
}
