package models

import (
	playlistModels "2021_1_Noskool_team/internal/app/playlists/models"
	"encoding/json"
	"errors"
)

type OtherUserFullInformation struct {
	UserID        int                        `json:"user_id"`
	Nickname      string                     `json:"nickname"`
	Photo         string                     `json:"photo"`
	ISubscribed   bool                       `json:"I_subscribed"`
	Subscriptions []*OtherUser               `json:"subscriptions"`
	Subscribers   []*OtherUser               `json:"subscribers"`
	Playlists     []*playlistModels.Playlist `json:"playlists"`
}

type OtherUser struct {
	UserID      int    `json:"user_id"`
	Nickname    string `json:"nickname"`
	Photo       string `json:"photo"`
	ISubscribed bool   `json:"I_subscribed"`
}

func MarshalOtherUserFullInformation(data interface{}) ([]byte, error) {
	otherUser, ok := data.(*OtherUserFullInformation)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	body, err := otherUser.MarshalJSON()
	return body, err
}

func MarshalOtherUser(data interface{}) ([]byte, error) {
	otherUser, ok := data.(*OtherUser)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	body, err := otherUser.MarshalJSON()
	return body, err
}

func MarshalOtherUsers(data interface{}) ([]byte, error) {
	otherUsers, ok := data.([]*OtherUser)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	body, err := json.Marshal(otherUsers)
	return body, err
}
