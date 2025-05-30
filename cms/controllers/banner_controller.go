package controllers

import (
	"net/http"

	"cqrs_demo/database"
	"cqrs_demo/models"

	"github.com/gin-gonic/gin"
)

// Get all banners
func GetBanners(c *gin.Context) {
	var banners []models.Banner
	result := database.GetDB().Preload("Location").Find(&banners)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": banners})
}

// Get banner by ID
func GetBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	result := database.GetDB().Preload("Location").First(&banner, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banner not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": banner})
}

// Create banner
func CreateBanner(c *gin.Context) {
	var banner models.Banner

	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Create(&banner)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": banner})
}

// Update banner
func UpdateBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	if err := database.GetDB().First(&banner, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banner not found"})
		return
	}

	var updateData models.Banner
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Model(&banner).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": banner})
}

// Delete banner
func DeleteBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	result := database.GetDB().Delete(&banner, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banner not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banner deleted successfully"})
}
