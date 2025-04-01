package ginrestaurant

import (
	"net/http"
	"rest/common"
	"rest/component"
	"rest/modules/restaurant/restaurantbiz"
	"rest/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
)

func GetRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		result, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		result.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(&result))
	}
}