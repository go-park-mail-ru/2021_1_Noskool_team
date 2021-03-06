package usecase

import (
	"2021_1_Noskool_team/internal/microservices/auth"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"fmt"
)

const (
	oneDayTime = 86400
)

type SessionsUsecase struct {
	sessionsRepo auth.Repository
}

func NewSessionsUsecase(sessionRep auth.Repository) SessionsUsecase {
	return SessionsUsecase{
		sessionsRepo: sessionRep,
	}
}

func (usecase *SessionsUsecase) CheckSession(userID string) (*models.Sessions, error) {
	session := &models.Sessions{
		UserID:     userID,
		Expiration: oneDayTime,
	}

	session, err := usecase.sessionsRepo.CheckSession(session)
	if err != nil {
		fmt.Println(err)
		session.UserID = "-1"
		return session, err
	}

	return session, nil
}

func (usecase *SessionsUsecase) DeleteSession(userID string) error {
	session := &models.Sessions{
		UserID:     userID,
		Expiration: oneDayTime,
	}

	err := usecase.sessionsRepo.DeleteSession(session)
	if err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}

func (usecase *SessionsUsecase) CreateSession(userID string) (*models.Sessions, error) {
	session := &models.Sessions{
		UserID:     userID,
		Expiration: oneDayTime,
	}

	session, err := usecase.sessionsRepo.CreateSession(session)
	if err != nil {
		fmt.Println(err)
		session.UserID = "-1"
		return session, err
	}

	return session, nil
}
