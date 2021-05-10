package usecase

import (
	"2021_1_Noskool_team/internal/app/profiles"
	"2021_1_Noskool_team/internal/app/profiles/models"
)

type ProfilesUsecase struct {
	profilesRepo profiles.Repository
}

func NewProfilesUsecase(profilesRepo profiles.Repository) profiles.Usecase {
	return &ProfilesUsecase{
		profilesRepo: profilesRepo,
	}
}

func (usecase *ProfilesUsecase) Create(u *models.UserProfile) error {
	err := usecase.profilesRepo.Create(u)
	return err

}

func (usecase *ProfilesUsecase) Update(u *models.UserProfile) error {
	err := usecase.profilesRepo.Update(u)
	return err

}

func (usecase *ProfilesUsecase) FindByID(id string) (*models.UserProfile, error) {
	usr, err := usecase.profilesRepo.FindByID(id)
	return usr, err

}

func (usecase *ProfilesUsecase) UpdateAvatar(userID string, newAvatar string) {
	usecase.profilesRepo.UpdateAvatar(userID, newAvatar)
}

func (usecase *ProfilesUsecase) FindByLogin(nickname string) (*models.UserProfile, error) {
	usr, err := usecase.profilesRepo.FindByLogin(nickname)
	return usr, err
}
