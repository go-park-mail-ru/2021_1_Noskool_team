package usecase

import (
	"2021_1_Noskool_team/internal/app/album"
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	"2021_1_Noskool_team/internal/app/musicians"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/app/search/models"
	"2021_1_Noskool_team/internal/app/tracks"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
)

type SearchUsecase struct {
	tracksRep   tracks.Repository
	albumsRep   album.Repository
	musicianRep musicians.Repository
}

func NewSearchUsecase(trackRep tracks.Repository, albumRep album.Repository,
	musicianRep musicians.Repository) *SearchUsecase {
	return &SearchUsecase{
		tracksRep:   trackRep,
		albumsRep:   albumRep,
		musicianRep: musicianRep,
	}
}

func (usecase *SearchUsecase) SearchContent(searchQuery string) *models.Search {
	search := &models.Search{}
	tracks, _ := usecase.tracksRep.SearchTracks(searchQuery)
	if tracks == nil {
		search.Tracks = make([]*trackModels.Track, 0)
	} else {
		search.Tracks = tracks
	}

	albums, _ := usecase.albumsRep.SearchAlbums(searchQuery)
	if albums == nil {
		search.Albums = make([]*albumsModels.Album, 0)
	} else {
		search.Albums = albums
	}

	musicians, _ := usecase.musicianRep.SearchMusicians(searchQuery)
	if musicians == nil {
		search.Musicians = make([]*musiciansModels.Musician, 0)
	} else {
		search.Musicians = musicians
	}

	return search
}
