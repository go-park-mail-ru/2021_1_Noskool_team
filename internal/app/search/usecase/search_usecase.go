package usecase

import (
	"2021_1_Noskool_team/internal/app/album"
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	"2021_1_Noskool_team/internal/app/musicians"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/app/playlists"
	playlistsModels "2021_1_Noskool_team/internal/app/playlists/models"
	"2021_1_Noskool_team/internal/app/search/models"
	"2021_1_Noskool_team/internal/app/tracks"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
)

type SearchUsecase struct {
	tracksRep   tracks.Repository
	albumsRep   album.Repository
	musicianRep musicians.Repository
	playlistRep playlists.Repository
}

func NewSearchUsecase(trackRep tracks.Repository, albumRep album.Repository,
	musicianRep musicians.Repository, playlistRep playlists.Repository) *SearchUsecase {
	return &SearchUsecase{
		tracksRep:   trackRep,
		albumsRep:   albumRep,
		musicianRep: musicianRep,
		playlistRep: playlistRep,
	}
}



func (usecase *SearchUsecase) SearchContent(searchQuery string) *models.Search {
	search := &models.Search{}
	tracks, _ := usecase.tracksRep.SearchTracks(searchQuery)
	if tracks == nil {
		search.Tracks = make([]*models.TrackWithAlbum, 0)
	} else {
		search.Tracks = usecase.ConvertTracks(tracks)
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

	playlists, _ := usecase.playlistRep.SearchPlaylists(searchQuery)
	if playlists == nil {
		search.Playlists = make([]*playlistsModels.Playlist, 0)
	} else {
		search.Playlists = playlists
	}

	return search
}

func (usecase *SearchUsecase) ConvertTracks(tracks []*trackModels.Track) []*models.TrackWithAlbum{
	newTracks := make([]*models.TrackWithAlbum, len(tracks))

	for idx, track := range tracks {
		newTracks[idx] = models.ConvertTrackToTrackWithAlbum(track)
		albums, err := usecase.albumsRep.GetAlbumsByTrackID(track.TrackID)
		if err == nil && albums != nil && len(*albums) > 0 {
			newTracks[idx].Album = (*albums)[0].AlbumID
		}
	}
	return newTracks
}
