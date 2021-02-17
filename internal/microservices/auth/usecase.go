package auth

import "2021_1_Noskool_team/internal/microservices/auth/models"

type Usecase interface {
	CheckSession(string) (*models.Sessions, error)
	DeleteSession(string) error
	CreateSession(string) (*models.Sessions, error)
}
