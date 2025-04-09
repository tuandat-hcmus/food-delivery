package ginuser

import (
	"net/http"
	"rest/common"
	"rest/component"
	"rest/component/hasher"
	"rest/component/tokenprovider/jwt"
	"rest/modules/user/userbiz"
	"rest/modules/user/usermodel"
	"rest/modules/user/userstorage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())
		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, appCtx.NewTokenConfig())
		account, err := biz.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}