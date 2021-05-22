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
		acceptableMusicGenres := []string{
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
		reciveGenres, ok := value.([]string)
		if !ok {
			return errors.New("must be a valid json type")
		}
		if len(reciveGenres) == 0 {
			return nil
		}
		for _, Gengre := range reciveGenres {
			if !Contains(acceptableMusicGenres, Gengre) {
				return errors.New("must be a valid genre type")
			}
		}
		if !CheckAllUniq(reciveGenres) {
			return errors.New("genres should not be repeated")
		}
		return nil
	}
}

func Contains(set []string, x string) bool {
	for _, n := range set {
		if x == n {
			return true
		}
	}
	return false
}

func CheckAllUniq(set []string) bool {
	for idx, item := range set {
		for idy, curr := range set {
			if item == curr && idx != idy {
				return false
			}
		}
	}
	return true
}
