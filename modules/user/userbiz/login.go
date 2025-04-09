package userbiz

import (
	"context"
	"rest/component"
	"rest/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type TokenConfig interface {
	GetAtExp() int
	GetRtExp() int
}

type loginBiz struct {
	appCtx component.AppContext
	storeUser LoginStorage
	tokenProvider 
	haser Hasher
	tkCfg TokenConfig
}