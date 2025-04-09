package userbiz

import (
	"context"
	// "rest/component"
	"rest/component/tokenprovider"
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
	// appCtx component.AppContext
	storeUser LoginStorage
	tokenProvider tokenprovider.Provider
	hasher Hasher
	tkCfg TokenConfig
}

func NewLoginBiz (storeUser LoginStorage, tokenProvider tokenprovider.Provider, 
	hasher Hasher, tkCfg TokenConfig) *loginBiz {
		return &loginBiz{
			storeUser: storeUser,
			tokenProvider: tokenProvider,
			hasher: hasher,
			tkCfg: tkCfg,
		}
	}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, 
	}
}
