package middleware

import (
	"errors"
	"rest/common"
	"rest/component"
	"rest/component/tokenprovider/jwt"
	"rest/modules/user/userstorage"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err, 
		"wrong authen header", 
		"ErrWrongAuthenHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// get token from header 
// validate token and parse payload
// from user_id in payload, find the user in db
func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())
	return func (c *gin.Context)  {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		// find user from DB
		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlStore(db)
		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}
		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}