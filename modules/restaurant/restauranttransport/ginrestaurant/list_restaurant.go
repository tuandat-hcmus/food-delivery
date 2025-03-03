package ginrestaurant

import (
	"net/http"
	"rest/common"
	"rest/component"
	"rest/modules/restaurant/restaurantbiz"
	"rest/modules/restaurant/restaurantmodel"
	"rest/modules/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err !=  nil {
			c.JSON(401, gin.H {
				"error" : err.Error(),
			})
			return
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err !=  nil {
			c.JSON(401, gin.H {
				"error" : err.Error(),
			})
			return
		}
		paging.Fulfill()
		store := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, gin.H {
				"error" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}

}