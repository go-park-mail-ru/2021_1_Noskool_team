package auth

import "testWorkWithAuth/internal/microservices/auth/models"

type Repository interface {
	CreateSession(*models.Sessions) (*models.Sessions, error)
	CheckSession(*models.Sessions) (*models.Sessions, error)
	DeleteSession(*models.Sessions) error
}
