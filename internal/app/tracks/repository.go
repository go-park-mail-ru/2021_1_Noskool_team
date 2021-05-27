package tracks

import (
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/app/tracks/models"
	commonModels "2021_1_Noskool_team/internal/models"
)

type Repository interface {
	GetTrackByID(trackID int) (*models.Track, error)
	GetTracksByTittle(trackTittle string) ([]*models.Track, error)
	GetTrackByMusicianID(musicianID int) ([]*models.Track, error)
	CreateTrack(*models.Track) (*models.Track, error)
	UploadPicture(trackID int, audioPath string) error
	UploadAudio(trackID int, audioPath string) error
	GetTracksByUserID(userID int) ([]*models.Track, error)
	GetFavoriteTracks(userID int, pagination *commonModels.Pagination) ([]*models.Track, error)
	AddTrackToFavorites(userID, trackID int) error
	DeleteTrackFromFavorites(userID, trackID int) error
	GetTracksByAlbumID(albumID int) ([]*models.Track, error)
	GetTracksByGenreID(genreID int) ([]*models.Track, error)
	AddTrackToMediateka(userID, trackID int) error
	DeleteTrackFromMediateka(userID, trackID int) error
	SearchTracks(searchQuery string) ([]*models.Track, error)
	GetTop20Tracks() ([]*models.Track, error)
	GetBillbordTopCharts() ([]*models.Track, error)
	GetHistory(userID int) ([]*models.Track, error)
	AddToHistory(userID, trackID int) error
	GetMusiciansGenresAndAlbums(tracks []*models.Track) []*models.Track
	GetMusicianByTrackID(trackID int) []*musiciansModels.Musician
	GetAlbumsByTrackID(trackID int) []*albumsModels.Album
	GetGenreByTrackID(trackID int) []*commonModels.Genre
	GetTopTrack() ([]*models.Track, error)

	CheckTrackInMediateka(userID, trackID int) error
	CheckTrackInFavorite(userID, trackID int) error

	IncrementLikes(trackID int) error
	DecrementLikes(trackID int) error
}
