package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minsecretkeysize = 32

type JWTmaker struct {
	Secret string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minsecretkeysize {
		fmt.Printf("the secret key size is lessthan 32 character")
	}
	return &JWTmaker{secretKey}, nil
}

func (maker *JWTmaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.Secret))
}

func (maker *JWTmaker) VerifyToken(token string) (*Payload, error) {
	KeyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.Secret), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, KeyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken

	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
