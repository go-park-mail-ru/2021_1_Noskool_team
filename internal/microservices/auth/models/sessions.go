package models

type Sessions struct {
	UserID     string
	Hash       string
	Expiration int
}

type Result struct {
	ID     string
	Hash   string
	Status string
}
