package ginrestaurant

import (
	"net/http"
	"rest/common"
	"rest/component"
	"rest/modules/restaurant/restaurantbiz"
	"rest/modules/restaurant/restaurantmodel"
	"rest/modules/restaurant/restaurantstorage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, gin.H {
				"error" : err.Error(),
			})
			return
		}

		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H {
				"error" : err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		
		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(401, gin.H {
				"error" : err.Error(), 
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}