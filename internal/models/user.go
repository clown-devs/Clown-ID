package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Validate user public information
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

// Check public fields and password
func (u *User) BeforeCreate() error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.CheckPassword(); err != nil {
		return err
	}

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
		return fmt.Errorf("password: is empty")
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("password: length is less than 6")
	}

	return nil
}
