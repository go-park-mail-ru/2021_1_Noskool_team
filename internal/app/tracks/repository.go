package tracks

import (
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
	CheckTrackInMediateka(userID, trackID int) error
}
