package profiles

import "2021_1_Noskool_team/internal/app/profiles/models"

type Usecase interface {
	Create(u *models.UserProfile) error
	Update(u *models.UserProfile, withPassword bool) error
	FindByID(id string) (*models.UserProfile, error)
	UpdateAvatar(userID string, newAvatar string)
	FindByLogin(nickname string) (*models.UserProfile, error)
}
