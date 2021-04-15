package models

type Response struct {
	Error HTTPError   `json:"error"`
	Body  interface{} `json:"body"`
}
