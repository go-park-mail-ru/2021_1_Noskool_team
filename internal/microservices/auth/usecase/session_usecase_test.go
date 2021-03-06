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
		UserID:     "1",
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().CheckSession(gomock.Eq(expectedSession)).
		Return(expectedSession, nil)

	track, err := mockUsecase.CheckSession(expectedSession.UserID)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedSession, track)
}

func TestCheckSessionFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_auth.NewMockRepository(ctrl)
	mockUsecase := NewSessionsUsecase(mockRepo)

	expectedSession := &models.Sessions{
		UserID:     "-1",
		Expiration: 86400,
	}

	mockRepo.
		EXPECT().CheckSession(gomock.Eq(expectedSession)).
		Return(expectedSession, errors.New("session bd fail"))

	track, err := mockUsecase.CheckSession(expectedSession.UserID)
	assert.Equal(t, err, errors.New("session bd fail"))
	assert.Equal(t, expectedSession, track)
}
