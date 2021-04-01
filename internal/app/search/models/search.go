package models

import (
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
)

type Search struct {
	Tracks    []*trackModels.Track        `json:"tracks,omitempty"`
	Albums    []*albumsModels.Album       `json:"albums,omitempty"`
	Musicians []*musiciansModels.Musician `json:"musicians,omitempty"`
}
