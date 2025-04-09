package userbiz

import (
	"context"
	"fmt"
	// "rest/component"
	"rest/common"
	"rest/component/tokenprovider"
	"rest/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	// appCtx component.AppContext
	storeUser LoginStorage
	tokenProvider tokenprovider.Provider
	hasher Hasher
	tkCfg common.TokenConfig
}

func NewLoginBiz (storeUser LoginStorage, tokenProvider tokenprovider.Provider, 
	hasher Hasher, tkCfg common.TokenConfig) *loginBiz {
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
		return nil, usermodel.ErrEmailExisted
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	fmt.Println(passHashed)
	fmt.Println(user.Password)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayLoad {
		UserId: user.Id, 
		Role: user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetAtExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetRtExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return usermodel.NewAccount(accessToken, refreshToken), nil
}
