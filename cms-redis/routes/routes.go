// routes/routes.go - Định nghĩa routes cho CMS Redis

package routes

import (
	"cms_redis/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/locations", controllers.GetLocations)
		v1.GET("/locations/:id", controllers.GetLocation)
		v1.POST("/locations", controllers.CreateLocation)
		v1.PUT("/locations/:id", controllers.UpdateLocation)
		v1.DELETE("/locations/:id", controllers.DeleteLocation)
	}
}