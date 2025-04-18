package jwt

import (
	"fmt"
	"rest/component/tokenprovider"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtProvider struct {
	secret string
}

func NewTokenJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	PayLoad tokenprovider.TokenPayLoad `json:"payload"`
	jwt.RegisteredClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayLoad, expiry int) (*tokenprovider.Token, error) {
	// generate jwt
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Second * time.Duration(expiry))),
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		},
	})
	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}
	// return token 
	return &tokenprovider.Token{
		Token: myToken,
		Expiry: expiry,
		Created: time.Now().UTC(),
	}, nil
}

// validate + parse payload
func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayLoad, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	fmt.Print(err)

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}
	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}
	return &claims.PayLoad, nil
}