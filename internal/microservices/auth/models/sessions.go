package models

type Sessions struct {
	UserID     string
	Expiration int
}

type Result struct {
	ID     int
	Status string
}
