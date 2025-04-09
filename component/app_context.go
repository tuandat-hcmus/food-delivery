package component

import (
	"rest/common"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	NewTokenConfig() common.TokenConfig
}

type tokenConfig struct {
	atExp int
	rtExp int
}

func NewTokenConfig(atExp, rtExp int) *tokenConfig {
	return &tokenConfig{atExp: atExp, rtExp: rtExp}
}

func (t *tokenConfig) GetAtExp() int {
	return t.atExp
}

func (t *tokenConfig) GetRtExp() int {
	return t.rtExp
}

type appCtx struct {
	db *gorm.DB
	secretKey string
	atExp int
	rtExp int
}

func NewAppContext(db *gorm.DB, secretKey string, atExp, rtExp int) *appCtx {
	return &appCtx{
		db: db, 
		secretKey: secretKey,
		atExp: atExp,
		rtExp: rtExp,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) NewTokenConfig() common.TokenConfig {
	return NewTokenConfig(ctx.atExp, ctx.rtExp)
}

