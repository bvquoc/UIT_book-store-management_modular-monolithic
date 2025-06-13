package booklocationtransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/booklocation/booklocationbiz"
	"book-store-management-backend/module/booklocation/booklocationrepo"
	"book-store-management-backend/module/booklocation/booklocationstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllBookLocation
// @BasePath /v1
// @Security BearerAuth
// @Summary Get all book locations
// @Tags book-locations
// @Accept json
// @Produce json
// @Param active query bool false "Just get active locations"
// @Response 200 {object} common.ResSuccess "list book locations"
// @Response 400 {object} common.AppError "error"
// @Router /book-locations [get]
func GetAllBookLocation(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the "active" filter from the query parameters
		activeParam := c.DefaultQuery("active", "true")
		justGetAllActiveLocation := activeParam == "true"

		// Initialize database connection
		db := appCtx.GetMainDBConnection()
		bookLocationStore := booklocationstore.NewSQLStore(db)
		bookLocationRepo := booklocationrepo.NewGetAllBookLocationRepo(bookLocationStore)
		bookLocationBiz := booklocationbiz.NewGetAllBookLocationBiz(bookLocationRepo)

		// Call the business logic to fetch the book locations
		bookLocations, err := bookLocationBiz.GetAllBookLocation(c.Request.Context(), justGetAllActiveLocation)
		if err != nil {
			panic(err)
		}

		// Return the results
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(bookLocations))
	}
}
