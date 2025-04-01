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

func UpdateRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		
		if err := biz.UpdateRestaurant(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}