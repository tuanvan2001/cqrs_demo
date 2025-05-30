package controllers

import (
	"net/http"

	"cqrs_demo/database"
	"cqrs_demo/models"

	"github.com/gin-gonic/gin"
)

// Get all games
func GetGames(c *gin.Context) {
	var games []models.Game
	result := database.GetDB().Preload("Location").Find(&games)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// Get game by ID
func GetGame(c *gin.Context) {
	id := c.Param("id")
	var game models.Game

	result := database.GetDB().Preload("Location").First(&game, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// Create game
func CreateGame(c *gin.Context) {
	var game models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Create(&game)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": game})
}

// Update game
func UpdateGame(c *gin.Context) {
	id := c.Param("id")
	var game models.Game

	if err := database.GetDB().First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	var updateData models.Game
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Model(&game).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// Delete game
func DeleteGame(c *gin.Context) {
	id := c.Param("id")
	var game models.Game

	result := database.GetDB().Delete(&game, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}
