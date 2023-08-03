package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.ID, validation.Required),
	); err != nil {
		return err
	}

	err := u.CheckPassword()
	return err
}

func (u *User) BeforeCreate() error {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.EncryptedPassword = string(encrypted)
	return err
}

func (u *User) Sanitize() {
	u.Password = ""
}

// returns nil if password is ok. Else - error
func (u *User) CheckPassword() error {
	if u.Password == "" {
		return fmt.Errorf("password is empty")
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("password len is less than 6")
	}

	return nil
}
