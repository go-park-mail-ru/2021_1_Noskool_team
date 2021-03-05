package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// UserProfile ...
type UserProfile struct {
	ProfileID         int    `json:"-"`
	Email             string `json:"email"`
	Login             string `json:"login"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Validate ....
func (u *UserProfile) Validate(withPassword bool) error {
	if withPassword {
		return validation.ValidateStruct(
			u,
			validation.Field(&u.Email, validation.Required, is.Email),
			validation.Field(&u.Login, validation.Required, validation.Length(6, 64)),
			validation.Field(&u.Password, validation.By(requiredIF(u.EncryptedPassword == "")), validation.Length(6, 32)),
		)
	}
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Login, validation.Required, validation.Length(6, 64)),
	)

}

// BeforeCreate ...
func (u *UserProfile) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

// Sanitize ...
func (u *UserProfile) Sanitize() {
	u.Password = ""
}

func (u *UserProfile) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
