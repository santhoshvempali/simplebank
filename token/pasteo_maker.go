package token

import (
	"time"

	"github.com/aead/chacha20poly1305"

	"github.com/o1egl/paseto"
)

type PasteoMaker struct {
	pasteo      *paseto.V2
	symetricKey []byte
}

func NewPasteoMaker(symetricKey string) (Maker, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, paseto.ErrTokenValidationError
	}
	maker := &PasteoMaker{
		pasteo:      paseto.NewV2(),
		symetricKey: []byte(symetricKey),
	}
	return maker, nil
}

func (maker *PasteoMaker) CreateToken(username string, duraton time.Duration) (string, error) {
	payload, err := NewPayload(username, duraton)

	if err != nil {
		return "", err
	}
	return maker.pasteo.Encrypt(maker.symetricKey, payload, nil)
}
func (maker *PasteoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.pasteo.Decrypt(token, maker.symetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil

}
