package token

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
)

var (
	ErrorExpiredToken = errors.New("token is expired")
	ErrorInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"exoried_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrorExpiredToken
	}
	return nil
}
