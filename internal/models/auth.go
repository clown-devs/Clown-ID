package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
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

type Application struct {
	ID   string `json:"id" example:"1"`
	Name string `json:"name" example:"clown-space"`
}

type Client struct {
	ID   string `json:"id" example:"1"`
	Name string `json:"name" example:"android"`
}

func (app *Application) Validate() error {
	return validation.ValidateStruct(app,
		validation.Field(&app.Name, validation.Required),
	)
}

func (client *Client) Validate() error {
	return validation.ValidateStruct(client,
		validation.Field(&client.Name, validation.Required),
	)
}
