package booklocationtransport

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	// Grouping the routes under "/book-locations"
	bookLocations := router.Group("/book-locations", middleware.RequireAuth(appCtx))
	{
		// Create a new book location
		bookLocations.POST("", CreateBookLocation(appCtx))
	}

	// Get all book locations
	bookLocations.GET("/", GetAllBookLocation(appCtx))

}
