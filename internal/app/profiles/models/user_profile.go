package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

//easyjson:json
type UserProfile struct {
	ProfileID         int      `json:"user_id"`
	Email             string   `json:"email"`
	Login             string   `json:"login"`
	Name              string   `json:"first_name"`
	Surname           string   `json:"second_name"`
	Password          string   `json:"password,omitempty"`
	EncryptedPassword string   `json:"-"`
	Avatar            string   `json:"avatar"`
	FavoriteGenre     []string `json:"favorite_genre"`
}

// Validate ....
func (u *UserProfile) ValidateForCreate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email,
			validation.Required,
			is.Email.Error("Некорректный email")),
		validation.Field(&u.Login,
			validation.Required,
			validation.Length(6, 64).Error("Логин должен содержать от 6 до 64 символов")),
		validation.Field(&u.Password,
			validation.By(requiredIF(u.EncryptedPassword == "")),
			validation.Length(6, 32).Error("Пароль должен содержать от 6 до 32 символов")),
	)
}

func (u *UserProfile) ValidationForUpdate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email,
			validation.By(requiredIF(u.Email != "")),
			is.Email.Error("Некорректный email")),
		validation.Field(&u.Login,
			validation.By(requiredIF(u.Login != "")),
			validation.Length(6, 64).Error("Логин должен содержать от 6 до 64 символов")),
		validation.Field(&u.Name,
			validation.By(requiredIF(u.Name != "")),
			validation.Length(2, 64).Error("Имя должно содержать от 2 до 64 символов")),
		validation.Field(&u.Surname,
			validation.By(requiredIF(u.Surname != "")),
			validation.Length(2, 128).Error("Фамилия должна содержать от 2 до 128 символов")),
		validation.Field(&u.FavoriteGenre, validation.By(chrckGenres())),
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
