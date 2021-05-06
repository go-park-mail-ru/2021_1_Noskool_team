package models

import "github.com/microcosm-cc/bluemonday"

//easyjson:json
type RequestLogin struct {
	Login    string `json:"nickname"`
	Password string `json:"password"`
}

//easyjson:json
type ProfileRequest struct {
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Nickname      string   `json:"nickname"`
	Name          string   `json:"first_name"`
	Surname       string   `json:"second_name"`
	FavoriteGenre []string `json:"favorite_genre"`
}

type ProfileForRegistrate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func (req *RequestLogin) Sanitize(sanitizer *bluemonday.Policy) {
	req.Login = sanitizer.Sanitize(req.Login)
	req.Password = sanitizer.Sanitize(req.Password)
}

func (req *ProfileRequest) Sanitize(sanitizer *bluemonday.Policy) {
	req.Email = sanitizer.Sanitize(req.Email)
	req.Password = sanitizer.Sanitize(req.Password)
	req.Nickname = sanitizer.Sanitize(req.Nickname)
	req.Name = sanitizer.Sanitize(req.Name)
	req.Surname = sanitizer.Sanitize(req.Surname)

	for i := range req.FavoriteGenre {
		req.FavoriteGenre[i] = sanitizer.Sanitize(req.FavoriteGenre[i])
	}
}

func (req *ProfileForRegistrate) Sanitize(sanitizer *bluemonday.Policy) {
	req.Email = sanitizer.Sanitize(req.Email)
	req.Password = sanitizer.Sanitize(req.Password)
	req.Nickname = sanitizer.Sanitize(req.Nickname)
}
