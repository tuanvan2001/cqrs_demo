package routes

import (
	"cqrs_demo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		// Location routes
		locations := api.Group("/locations")
		{
			locations.GET("", controllers.GetLocations)
			locations.GET("/:id", controllers.GetLocation)
			locations.POST("", controllers.CreateLocation)
			locations.PUT("/:id", controllers.UpdateLocation)
			locations.DELETE("/:id", controllers.DeleteLocation)
		}

		// Game routes
		games := api.Group("/games")
		{
			games.GET("", controllers.GetGames)
			games.GET("/:id", controllers.GetGame)
			games.POST("", controllers.CreateGame)
			games.PUT("/:id", controllers.UpdateGame)
			games.DELETE("/:id", controllers.DeleteGame)
		}

		// News routes
		news := api.Group("/news")
		{
			news.GET("", controllers.GetNews)
			news.GET("/:id", controllers.GetNewsItem)
			news.POST("", controllers.CreateNews)
			news.PUT("/:id", controllers.UpdateNews)
			news.DELETE("/:id", controllers.DeleteNews)
		}

		// Setting routes
		settings := api.Group("/settings")
		{
			settings.GET("", controllers.GetSettings)
			settings.GET("/:id", controllers.GetSetting)
			settings.POST("", controllers.CreateSetting)
			settings.PUT("/:id", controllers.UpdateSetting)
			settings.DELETE("/:id", controllers.DeleteSetting)
		}

		// Banner routes
		banners := api.Group("/banners")
		{
			banners.GET("", controllers.GetBanners)
			banners.GET("/:id", controllers.GetBanner)
			banners.POST("", controllers.CreateBanner)
			banners.PUT("/:id", controllers.UpdateBanner)
			banners.DELETE("/:id", controllers.DeleteBanner)
		}
	}
}
