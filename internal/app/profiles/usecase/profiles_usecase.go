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

func (usecase *ProfilesUsecase) UpdatePassword(id int, newPass string) error {
	err := usecase.profilesRepo.UpdatePassword(id, newPass)
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

func (usecase *ProfilesUsecase) SubscribeMeToSomebody(myID, otherUserID int) error {
	err := usecase.profilesRepo.SubscribeMeToSomebody(myID, otherUserID)
	return err
}

func (usecase *ProfilesUsecase) UnsubscribeMeToSomebody(myID, otherUserID int) error {
	err := usecase.profilesRepo.UnsubscribeMeToSomebody(myID, otherUserID)
	return err
}

func (usecase *ProfilesUsecase) GetOtherUserPage(myID, otherUserID int) (*models.OtherUserFullInformation, error) {
	otherUser, err := usecase.profilesRepo.GetOtherUserPage(otherUserID)
	if err != nil {
		return nil, err
	}
	otherUser.ISubscribed = usecase.profilesRepo.CheckIsMySubscriber(myID, otherUserID)
	otherUserFullInf := &models.OtherUserFullInformation{
		UserID:      otherUser.UserID,
		Nickname:    otherUser.Nickname,
		Photo:       otherUser.Photo,
		ISubscribed: otherUser.ISubscribed,
	}
	otherUserFullInf.Subscribers, _ = usecase.profilesRepo.GetSubscribers(otherUserFullInf.UserID)
	if otherUserFullInf.Subscribers != nil {
		for idx := 0; idx < len(otherUserFullInf.Subscribers); idx++ {
			otherUserFullInf.Subscribers[idx].ISubscribed = usecase.profilesRepo.CheckIsMySubscriber(myID,
				otherUserFullInf.Subscribers[idx].UserID)
		}
	}
	otherUserFullInf.Subscriptions, _ = usecase.profilesRepo.GetSubscriptions(otherUserFullInf.UserID)
	if otherUserFullInf.Subscriptions != nil {
		for idx := 0; idx < len(otherUserFullInf.Subscriptions); idx++ {
			otherUserFullInf.Subscriptions[idx].ISubscribed = usecase.profilesRepo.CheckIsMySubscriber(myID,
				otherUserFullInf.Subscriptions[idx].UserID)
		}
	}
	return otherUserFullInf, nil
}

func (usecase *ProfilesUsecase) SearchTracks(searchQuery string) ([]*models.OtherUser, error) {
	otherUsers, err := usecase.profilesRepo.SearchTracks(searchQuery)
	return otherUsers, err
}
