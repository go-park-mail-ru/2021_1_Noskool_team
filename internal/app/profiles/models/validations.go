package models

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

func requiredIF(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}

func chrckGenres() validation.RuleFunc {
	return func(value interface{}) error {
		acceptableMusicGenres := [12]string{
			"classical",
			"jazz",
			"rap",
			"electronic",
			"rock",
			"disco",
			"fusion",
			"pop",
			"country",
			"blues",
			"reggae",
			"indie",
		}
		reciveGenre, ok := value.(string)
		if !ok {
			return errors.New("must be a valid genre type")
		}
		if reciveGenre == "" {
			return errors.New("cannot be blank")
		}
		for _, acceptableGenre := range acceptableMusicGenres {
			if reciveGenre == acceptableGenre {
				return nil
			}
		}
		return errors.New("must be a valid genre type")
	}
}
