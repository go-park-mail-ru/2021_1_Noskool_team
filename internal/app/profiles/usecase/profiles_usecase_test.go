package usecase

import (
	mockProfiles "2021_1_Noskool_team/internal/app/profiles/mocks"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockProfiles.NewMockRepository(ctrl)
	mockUsecase := NewProfilesUsecase(mockRepo)

	expectedUser := models.UserProfile{
		ProfileID:     1,
		Email:         "test222@gmail.com",
		Login:         "test222",
		Name:          "Name222",
		Surname:       "Surname222",
		Password:      "Password222",
		Avatar:        "Avatar222",
		FavoriteGenre: []string{"rock"},
	}

	mockRepo.EXPECT().FindByID("1").Return(&expectedUser, nil)

	user, err := mockUsecase.FindByID("1")
	assert.Equal(t, err, nil)
	assert.Equal(t, &expectedUser, user)
}

func TestFindByLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockProfiles.NewMockRepository(ctrl)
	mockUsecase := NewProfilesUsecase(mockRepo)

	expectedUser := models.UserProfile{
		ProfileID:     1,
		Email:         "test222@gmail.com",
		Login:         "test222",
		Name:          "Name222",
		Surname:       "Surname222",
		Password:      "Password222",
		Avatar:        "Avatar222",
		FavoriteGenre: []string{"rock"},
	}

	mockRepo.EXPECT().FindByLogin("test222").Return(&expectedUser, nil)

	user, err := mockUsecase.FindByLogin("test222")
	assert.Equal(t, err, nil)
	assert.Equal(t, &expectedUser, user)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockProfiles.NewMockRepository(ctrl)
	mockUsecase := NewProfilesUsecase(mockRepo)

	User := models.UserProfile{
		ProfileID:     1,
		Email:         "test222@gmail.com",
		Login:         "test222",
		Name:          "Name222",
		Surname:       "Surname222",
		Password:      "Password222",
		Avatar:        "Avatar222",
		FavoriteGenre: []string{"rock"},
	}

	mockRepo.EXPECT().Create(&User).Return(nil)

	err := mockUsecase.Create(&User)
	assert.Equal(t, err, nil)
}

func TestUpdateTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockProfiles.NewMockRepository(ctrl)
	mockUsecase := NewProfilesUsecase(mockRepo)

	User := models.UserProfile{
		ProfileID:     1,
		Email:         "test222@gmail.com",
		Login:         "test222",
		Name:          "Name222",
		Surname:       "Surname222",
		Password:      "Password222",
		Avatar:        "Avatar222",
		FavoriteGenre: []string{"rock"},
	}

	mockRepo.EXPECT().Update(&User, true).Return(nil)

	err := mockUsecase.Update(&User, true)
	assert.Equal(t, err, nil)
}

func TestUpdateFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockProfiles.NewMockRepository(ctrl)
	mockUsecase := NewProfilesUsecase(mockRepo)

	User := models.UserProfile{
		ProfileID:     1,
		Email:         "test222@gmail.com",
		Login:         "test222",
		Name:          "Name222",
		Surname:       "Surname222",
		Password:      "Password222",
		Avatar:        "Avatar222",
		FavoriteGenre: []string{"rock"},
	}

	mockRepo.EXPECT().Update(&User, false).Return(nil)

	err := mockUsecase.Update(&User, false)
	assert.Equal(t, err, nil)
}

func TestUpdateAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockProfiles.NewMockRepository(ctrl)
	mockUsecase := NewProfilesUsecase(mockRepo)

	mockRepo.EXPECT().UpdateAvatar("1", "path")

	mockUsecase.UpdateAvatar("1", "path")
}
