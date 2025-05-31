// controllers/location_controller.go - CRUD Location với RedisJSON/Redisearch

package controllers

import (
	"net/http"
	"strconv"
	"time"

	"cms_redis/database"
	"cms_redis/models"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/locations - Lấy tất cả location (Redisearch)
func GetLocations(c *gin.Context) {
	ctx, cancel := database.GetRedisContext()
	defer cancel()

	res, err := database.Rdb.Do(ctx,
		"FT.SEARCH", "idx:location", "*", "RETURN", "6", "id", "name", "address", "status", "created_at", "updated_at").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

// GET /api/v1/locations/:id - Lấy location theo id
func GetLocation(c *gin.Context) {
	id := c.Param("id")
	key := "location:" + id

	ctx, cancel := database.GetRedisContext()
	defer cancel()

	val, err := database.Rdb.Do(ctx, "JSON.GET", key).Result()
	if err != nil || val == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(val.(string)))
}

// POST /api/v1/locations - Tạo location mới
func CreateLocation(c *gin.Context) {
	var loc models.Location
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := database.GetRedisContext()
	defer cancel()

	id, err := database.Rdb.Incr(ctx, "location:id").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	loc.ID = uint(id)
	loc.CreatedAt = time.Now()
	loc.UpdatedAt = loc.CreatedAt
	key := "location:" + strconv.FormatUint(uint64(loc.ID), 10)
	locBytes, err := json.Marshal(loc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = database.Rdb.Do(ctx, "JSON.SET", key, "$", string(locBytes)).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": loc})
}

// PUT /api/v1/locations/:id - Cập nhật location
func UpdateLocation(c *gin.Context) {
	id := c.Param("id")
	key := "location:" + id
	var loc models.Location
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := database.GetRedisContext()
	defer cancel()

	parsedID, _ := strconv.ParseUint(id, 10, 64)
	loc.ID = uint(parsedID)
	loc.UpdatedAt = time.Now()
	locBytes, err := json.Marshal(loc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = database.Rdb.Do(ctx, "JSON.SET", key, "$", string(locBytes)).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": loc})
}

// DELETE /api/v1/locations/:id - Xóa location
func DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	key := "location:" + id

	ctx, cancel := database.GetRedisContext()
	defer cancel()

	_, err := database.Rdb.Del(ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
