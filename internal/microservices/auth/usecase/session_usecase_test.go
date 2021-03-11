package usecase

import (
	mock_auth "2021_1_Noskool_team/internal/microservices/auth/mocks"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_auth.NewMockRepository(ctrl)
	mockUsecase := NewSessionsUsecase(mockRepo)

	expectedSession := &models.Sessions{
		Hash:       "some hash",
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().CheckSession(gomock.Eq(expectedSession)).
		Return(expectedSession, nil)

	session, err := mockUsecase.CheckSession(expectedSession.Hash)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedSession, session)
}

func TestCheckSessionFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_auth.NewMockRepository(ctrl)
	mockUsecase := NewSessionsUsecase(mockRepo)

	expectedSession := &models.Sessions{
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().CheckSession(gomock.Eq(expectedSession)).
		Return(expectedSession, errors.New("session bd fail"))

	session, err := mockUsecase.CheckSession(expectedSession.UserID)
	assert.Equal(t, err, errors.New("session bd fail"))
	assert.Equal(t, expectedSession, session)
}

func TestDeleteSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_auth.NewMockRepository(ctrl)
	mockUsecase := NewSessionsUsecase(mockRepo)

	expectedSession := &models.Sessions{
		UserID:     "1",
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().DeleteSession(gomock.Eq(expectedSession)).
		Return(nil)

	err := mockUsecase.DeleteSession(expectedSession.UserID)
	assert.Equal(t, err, nil)
}

func TestDeleteSessionFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_auth.NewMockRepository(ctrl)
	mockUsecase := NewSessionsUsecase(mockRepo)

	expectedSession := &models.Sessions{
		UserID:     "1",
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().DeleteSession(gomock.Eq(expectedSession)).
		Return(errors.New("error in redis db"))

	err := mockUsecase.DeleteSession(expectedSession.UserID)
	assert.Equal(t, err, errors.New("error in redis db"))
}

func TestCreateSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_auth.NewMockRepository(ctrl)
	mockUsecase := NewSessionsUsecase(mockRepo)

	expectedSession := &models.Sessions{
		UserID:     "1",
		Hash:       "c4ca4238a0b923820dcc509a6f75849b",
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().CreateSession(gomock.Eq(expectedSession)).
		Return(expectedSession, nil)

	session, err := mockUsecase.CreateSession(expectedSession.UserID)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedSession, session)
}

func TestCreateSessionFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_auth.NewMockRepository(ctrl)
	mockUsecase := NewSessionsUsecase(mockRepo)

	expectedSession := &models.Sessions{
		UserID:     "1",
		Hash:       "c4ca4238a0b923820dcc509a6f75849b",
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().CreateSession(gomock.Eq(expectedSession)).
		Return(expectedSession, errors.New("error creating session"))

	session, err := mockUsecase.CreateSession(expectedSession.UserID)
	assert.Equal(t, err, errors.New("error creating session"))
	assert.Equal(t, expectedSession, session)
}
