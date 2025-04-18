package tokenprovider

import (
	"errors"
	"rest/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayLoad, expiry int) (*Token, error)
	Validate(token string) (*TokenPayLoad, error)
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayLoad struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"), 
		"token not found", 
		"ErrNotFound",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"), 
		"error encoding the token", 
		"ErrEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"), 
		"invalid token provided", 
		"ErrINvalidToken",
	)
)
