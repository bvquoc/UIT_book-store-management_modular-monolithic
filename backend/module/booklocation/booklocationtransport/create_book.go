package booklocationtransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booklocation/booklocationbiz"
	"book-store-management-backend/module/booklocation/booklocationmodel"
	"book-store-management-backend/module/booklocation/booklocationrepo"
	"book-store-management-backend/module/booklocation/booklocationstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBookLocation
// @BasePath /v1
// @Security BearerAuth
// @Summary Create Book Location
// @Tags book-locations
// @Accept json
// @Produce json
// @Param bookLocation body booklocationmodel.ReqCreateBookLocation true "Create Book Location"
// @Response 200 {object} common.ResSuccess "Created book location"
// @Router /book-locations [post]
func CreateBookLocation(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData booklocationmodel.ReqCreateBookLocation
		if err := c.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Starting a new database transaction
		db := appCtx.GetMainDBConnection().Begin()

		// Creating the necessary stores and dependencies
		gen := generator.NewShortIdGenerator()
		bookLocationStore := booklocationstore.NewSQLStore(db)
		bookLocationRepo := booklocationrepo.NewCreateBookLocationRepo(bookLocationStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		// Creating the business logic handler
		biz := booklocationbiz.NewCreateBookLocationBiz(gen, bookLocationRepo, requester)

		var resData booklocationmodel.ResCreateBookLocation
		// Business logic to create the book location
		if err := biz.CreateBookLocation(c.Request.Context(), &reqData, &resData); err != nil {
			db.Rollback()
			panic(err)
		}

		// Commit the transaction if successful
		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		// Respond with success
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(resData))
	}
}
