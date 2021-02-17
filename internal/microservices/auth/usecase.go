package auth

import "2021_1_Noskool_team/internal/microservices/auth/models"

type Usecase interface {
	CheckSession(userID int) (*models.Sessions, error)
	DeleteSession(int) error
	CreateSession(int) (*models.Sessions, error)
}
