package userbiz

import (
	"context"
	"rest/common"
	"rest/modules/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStorage RegisterStorage
	hasher Hasher
}

func NewRegisterBiz(registerStorage RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStorage: registerStorage,
		hasher: hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err == common.RecordNotFound {
		data.Salt = common.GenSalt(50)
		data.Password = biz.hasher.Hash(data.Password + data.Salt)
		data.Role = "user"
		data.Status = 1

		if e := biz.registerStorage.CreateUser(ctx, data); e != nil {
			return common.ErrCannotCreateEntity(usermodel.EntityName, e)
		}
		return nil
	}
	if user != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}
	return err
}