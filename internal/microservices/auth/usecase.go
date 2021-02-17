package auth

import "testWorkWithAuth/internal/microservices/auth/models"

type Usecase interface {
	CheckSession(userID int) (*models.Sessions, error)
	DeleteSession(int) error
	CreateSession(int) (*models.Sessions, error)
}
