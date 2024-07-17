package api

import (
	_ "api/api/docs"
	"api/api/handler"
	"api/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title User
// @version 1.0
// @description API Gateway
// @host localhost:8080
// BasePath: /
func Router(handler *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stories := router.Group("/api/v1/stories")
	stories.Use(middleware.Check)
	{
		stories.POST("", handler.CreateStories)
		stories.PUT("/:story_id", handler.UpdateStories)
		stories.DELETE("/:story_id", handler.DeleteStories)
		stories.GET("", handler.GetsStories)
		stories.GET("/:story_id", handler.GetStories)
		stories.POST("/:story_id/comments", handler.CommentStories)
		stories.GET("/:story_id/comments", handler.GetCommentStories)
		stories.POST("/:story_id/like", handler.LikeStories)
	}
	itineraries := router.Group("/api/v1/itineraries")
	itineraries.Use(middleware.Check)
	{
		itineraries.POST("", handler.CreateItineraries)
		itineraries.PUT("/:itinerary_id", handler.UpdateItineraries)
		itineraries.DELETE("/:itinerary_id", handler.DeleteItineraries)
		itineraries.GET("", handler.GetsItineraries)
		itineraries.GET("/:itinerary_id", handler.GetItineraries)
	}

	router.Use(middleware.Check)
	router.GET("/api/v1/destinations", handler.GetsDestinations)
	router.GET("/api/v1/destinations/:destination_id", handler.GetDestinations)
	router.POST("/api/v1/messages", handler.Message)
	router.GET("/api/v1/messages", handler.GetMessage)
	router.POST("/api/v1/travel-tips", handler.AddTravelTips)
	router.GET("/api/v1/travel-tips", handler.GetTravelTips)
	router.GET("/api/v1/users/:user_id/statistics", handler.StatisticsUser)
	router.GET("/api/v1/trending-destinations", handler.GetTrendingDestinations)

	user := router.Group("/api/v1/users")
	user.Use(middleware.Check)
	{
		user.GET("/profile", handler.GetUserProfile)
		user.PUT("/profile", handler.UpdateUserProfile)
		user.GET("", handler.ListUsersProfile)
		user.DELETE("/:user_id", handler.DeleteUserProfile)
		user.GET("/:user_id/activity", handler.ActivityUser)
		user.POST("/:user_id/follow", handler.FollowUser)
		user.GET("/:user_id/followers", handler.FollowersUsers)
	}


	return router
}
