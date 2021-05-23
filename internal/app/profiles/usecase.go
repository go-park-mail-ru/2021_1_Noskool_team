package profiles

import "2021_1_Noskool_team/internal/app/profiles/models"

type Usecase interface {
	Create(u *models.UserProfile) error
	Update(u *models.UserProfile) error
	FindByID(id string) (*models.UserProfile, error)
	UpdateAvatar(userID string, newAvatar string)
	FindByLogin(nickname string) (*models.UserProfile, error)
	UpdatePassword(int, string) error
	SubscribeMeToSomebody(myID, otherUserID int) error
	UnsubscribeMeToSomebody(myID, otherUserID int) error
	GetOtherUserPage(myID, otherUserID int) (*models.OtherUserFullInformation, error)
	SearchTracks(searchQuery string) ([]*models.OtherUser, error)
}
