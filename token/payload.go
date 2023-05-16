package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid Token")
)

type Payload struct {
	ID         uuid.UUID `json:id`
	Username   string    `json:username`
	Issued_at  time.Time `json:issued_at`
	Expires_at time.Time `json:expires_at`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:         tokenID,
		Username:   username,
		Issued_at:  time.Now(),
		Expires_at: time.Now().Add(duration),
	}
	return payload, err
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expires_at) {
		return ErrTokenExpired
	}
	return nil
}
