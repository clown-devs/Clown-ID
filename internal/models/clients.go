package models

import validation "github.com/go-ozzo/ozzo-validation"

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
