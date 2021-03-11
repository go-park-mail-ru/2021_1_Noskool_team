package auth

import "2021_1_Noskool_team/internal/microservices/auth/models"

type Repository interface {
	CreateSession(*models.Sessions) (*models.Sessions, error)
	CheckSession(*models.Sessions) (*models.Sessions, error)
	DeleteSession(*models.Sessions) error
}
