package usecase

import (
	"2021_1_Noskool_team/internal/app/search/models"
	"2021_1_Noskool_team/internal/app/tracks"
)

type SearchUsecase struct {
	tracksRep tracks.Repository
}

func NewSearchUsecase(trackRep tracks.Repository) *SearchUsecase {
	return &SearchUsecase{
		tracksRep: trackRep,
	}
}

func (usecase *SearchUsecase) SearchContent(searchQuery string) *models.Search {
	search := &models.Search{}
	tracks, _ := usecase.tracksRep.SearchTracks(searchQuery)
	search.Tracks = tracks

	return search
}
