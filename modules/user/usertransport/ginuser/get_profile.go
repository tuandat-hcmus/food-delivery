package ginuser

import (
	"net/http"
	"rest/common"
	"rest/component"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}