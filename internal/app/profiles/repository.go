package profiles

import "2021_1_Noskool_team/internal/app/profiles/models"

type Repository interface {
	Create(u *models.UserProfile) error
	Update(u *models.UserProfile) error
	FindByID(id string) (*models.UserProfile, error)
	UpdateAvatar(userID string, newAvatar string)
	FindByLogin(nickname string) (*models.UserProfile, error)
	UpdatePassword(int, string) error
	SubscribeMeToSomebody(myID, otherUserID int) error
	UnsubscribeMeToSomebody(myID, otherUserID int) error
	CheckIsMySubscriber(myID, otherUserID int) bool
	GetOtherUserPage(otherUserID int) (*models.OtherUser, error)
	GetSubscribers(userID int) ([]*models.OtherUser, error)
	GetSubscriptions(userID int) ([]*models.OtherUser, error)
	SearchTracks(searchQuery string) ([]*models.OtherUser, error)
}
