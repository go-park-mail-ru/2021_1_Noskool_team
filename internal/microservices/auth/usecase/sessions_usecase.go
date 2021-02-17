package usecase

import (
	"fmt"
	"testWorkWithAuth/internal/microservices/auth"
	"testWorkWithAuth/internal/microservices/auth/models"
	"testWorkWithAuth/internal/microservices/auth/repository"
)

type SessionsUsecase struct {
	sessionsRepo auth.Repository
}

func NewSessionsUsecase() SessionsUsecase {
	return SessionsUsecase{
		sessionsRepo: repository.NewSessionRepository("redis://user:@localhost:6379/0"),
	}
}

func (usecase *SessionsUsecase) CheckSession(userID int) (*models.Sessions, error) {
	session := &models.Sessions{
		UserID:     userID,
		Expiration: 5,
	}

	session, err := usecase.sessionsRepo.CheckSession(session)
	if err != nil {
		fmt.Println(err)
		session.UserID = -1
		return session, err
	}

	return session, nil
}

func (usecase *SessionsUsecase) DeleteSession(userID int) error {
	session := &models.Sessions{
		UserID:     userID,
		Expiration: 5,
	}

	err := usecase.sessionsRepo.DeleteSession(session)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (usecase *SessionsUsecase) CreateSession(userID int) (*models.Sessions, error) {
	session := &models.Sessions{
		UserID:     userID,
		Expiration: 5,
	}

	session, err := usecase.sessionsRepo.CreateSession(session)
	if err != nil {
		fmt.Println(err)
		session.UserID = -1
		return session, err
	}

	return session, nil
}
