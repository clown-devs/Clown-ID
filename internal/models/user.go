package models

import validation "github.com/go-ozzo/ozzo-validation"

type User struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.ID, validation.Required),
	)

	return err
}
