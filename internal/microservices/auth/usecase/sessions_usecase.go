package usecase

import (
	"2021_1_Noskool_team/internal/microservices/auth"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"crypto/md5"
	"encoding/hex"
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

func (usecase *SessionsUsecase) CheckSession(hash string) (*models.Sessions, error) {
	session := &models.Sessions{
		Hash:       hash,
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

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (usecase *SessionsUsecase) CreateSession(userID string) (*models.Sessions, error) {
	session := &models.Sessions{
		UserID:     userID,
		Hash:       "",
		Expiration: oneDayTime,
	}

	session.Hash = GetMD5Hash(session.UserID)
	fmt.Println(session.UserID)

	session, err := usecase.sessionsRepo.CreateSession(session)
	if err != nil {
		fmt.Println(err)
		session.UserID = "-1"
		return session, err
	}

	return session, nil
}
