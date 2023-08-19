package models

import (
	"time"
)

type RefreshToken struct {
	AppId     string `json:"app_id" example:"1"`
	ClientId  string `json:"client_id" example:"1"`
	UserId    string `json:"user_id" example:"1"`
	Token     string `json:"token" example:"550e8400-e29b-41d4-a716-446655440000"`
	ExpiresAt int64  `json:"expires_at" example:"1691407740"` // unixtime
}

func (token RefreshToken) IsExpired() bool {
	return token.ExpiresAt <= time.Now().Unix()
}
