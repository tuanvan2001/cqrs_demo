package controllers

import (
	"net/http"

	"cqrs_demo/database"
	"cqrs_demo/models"

	"github.com/gin-gonic/gin"
)

// Get all news
func GetNews(c *gin.Context) {
	var news []models.News
	result := database.GetDB().Find(&news)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

// Get news by ID
func GetNewsItem(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	result := database.GetDB().First(&news, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

// Create news
func CreateNews(c *gin.Context) {
	var news models.News

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Create(&news)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": news})
}

// Update news
func UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	if err := database.GetDB().First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	var updateData models.News
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Model(&news).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

// Delete news
func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	result := database.GetDB().Delete(&news, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}
