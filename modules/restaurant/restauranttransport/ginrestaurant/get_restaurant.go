package ginrestaurant

import (
	"net/http"
	"rest/common"
	"rest/component"
	"rest/modules/restaurant/restaurantbiz"
	"rest/modules/restaurant/restaurantstorage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		result, err := biz.GetRestaurant(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(&result))
	}
}