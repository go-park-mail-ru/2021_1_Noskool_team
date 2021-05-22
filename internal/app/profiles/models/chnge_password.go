package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/microcosm-cc/bluemonday"
)

type ChangePassword struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}

func (req *ChangePassword) Sanitize(sanitizer *bluemonday.Policy) {
	req.OldPassword = sanitizer.Sanitize(req.OldPassword)
	req.NewPassword = sanitizer.Sanitize(req.NewPassword)
}

func ValidateForChangePass(NewPassword string) error {
	return validation.Validate(NewPassword,
		validation.Required,
		validation.Length(6, 32).Error("Пароль должен содержать от 6 до 32 символов"))
}

func BeforeUpdatePass(NewPassword string) (string, error) {
	enc, err := encryptString(NewPassword)
	if err != nil {
		return "", err
	}
	return enc, nil
}
