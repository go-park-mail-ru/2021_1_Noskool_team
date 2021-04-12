// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_tracks is a generated GoMock package.
package mock_tracks

import (
	models "2021_1_Noskool_team/internal/app/tracks/models"
	models0 "2021_1_Noskool_team/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddToHistory mocks base method.
func (m *MockRepository) AddToHistory(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToHistory", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToHistory indicates an expected call of AddToHistory.
func (mr *MockRepositoryMockRecorder) AddToHistory(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToHistory", reflect.TypeOf((*MockRepository)(nil).AddToHistory), userID, trackID)
}

// AddTrackToFavorites mocks base method.
func (m *MockRepository) AddTrackToFavorites(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrackToFavorites", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTrackToFavorites indicates an expected call of AddTrackToFavorites.
func (mr *MockRepositoryMockRecorder) AddTrackToFavorites(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrackToFavorites", reflect.TypeOf((*MockRepository)(nil).AddTrackToFavorites), userID, trackID)
}

// AddTrackToMediateka mocks base method.
func (m *MockRepository) AddTrackToMediateka(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrackToMediateka", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTrackToMediateka indicates an expected call of AddTrackToMediateka.
func (mr *MockRepositoryMockRecorder) AddTrackToMediateka(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrackToMediateka", reflect.TypeOf((*MockRepository)(nil).AddTrackToMediateka), userID, trackID)
}

// CheckTrackInMediateka mocks base method.
func (m *MockRepository) CheckTrackInMediateka(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTrackInMediateka", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckTrackInMediateka indicates an expected call of CheckTrackInMediateka.
func (mr *MockRepositoryMockRecorder) CheckTrackInMediateka(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTrackInMediateka", reflect.TypeOf((*MockRepository)(nil).CheckTrackInMediateka), userID, trackID)
}

// CreateTrack mocks base method.
func (m *MockRepository) CreateTrack(arg0 *models.Track) (*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrack", arg0)
	ret0, _ := ret[0].(*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrack indicates an expected call of CreateTrack.
func (mr *MockRepositoryMockRecorder) CreateTrack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrack", reflect.TypeOf((*MockRepository)(nil).CreateTrack), arg0)
}

// DeleteTrackFromFavorites mocks base method.
func (m *MockRepository) DeleteTrackFromFavorites(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrackFromFavorites", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrackFromFavorites indicates an expected call of DeleteTrackFromFavorites.
func (mr *MockRepositoryMockRecorder) DeleteTrackFromFavorites(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrackFromFavorites", reflect.TypeOf((*MockRepository)(nil).DeleteTrackFromFavorites), userID, trackID)
}

// DeleteTrackFromMediateka mocks base method.
func (m *MockRepository) DeleteTrackFromMediateka(userID, trackID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrackFromMediateka", userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrackFromMediateka indicates an expected call of DeleteTrackFromMediateka.
func (mr *MockRepositoryMockRecorder) DeleteTrackFromMediateka(userID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrackFromMediateka", reflect.TypeOf((*MockRepository)(nil).DeleteTrackFromMediateka), userID, trackID)
}

// GetBillbordTopCharts mocks base method.
func (m *MockRepository) GetBillbordTopCharts() ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillbordTopCharts")
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillbordTopCharts indicates an expected call of GetBillbordTopCharts.
func (mr *MockRepositoryMockRecorder) GetBillbordTopCharts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillbordTopCharts", reflect.TypeOf((*MockRepository)(nil).GetBillbordTopCharts))
}

// GetFavoriteTracks mocks base method.
func (m *MockRepository) GetFavoriteTracks(userID int, pagination *models0.Pagination) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteTracks", userID, pagination)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteTracks indicates an expected call of GetFavoriteTracks.
func (mr *MockRepositoryMockRecorder) GetFavoriteTracks(userID, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteTracks", reflect.TypeOf((*MockRepository)(nil).GetFavoriteTracks), userID, pagination)
}

// GetHistory mocks base method.
func (m *MockRepository) GetHistory(userID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", userID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockRepositoryMockRecorder) GetHistory(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockRepository)(nil).GetHistory), userID)
}

// GetTop20Tracks mocks base method.
func (m *MockRepository) GetTop20Tracks() ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTop20Tracks")
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTop20Tracks indicates an expected call of GetTop20Tracks.
func (mr *MockRepositoryMockRecorder) GetTop20Tracks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTop20Tracks", reflect.TypeOf((*MockRepository)(nil).GetTop20Tracks))
}

// GetTrackByID mocks base method.
func (m *MockRepository) GetTrackByID(trackID int) (*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrackByID", trackID)
	ret0, _ := ret[0].(*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrackByID indicates an expected call of GetTrackByID.
func (mr *MockRepositoryMockRecorder) GetTrackByID(trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrackByID", reflect.TypeOf((*MockRepository)(nil).GetTrackByID), trackID)
}

// GetTrackByMusicianID mocks base method.
func (m *MockRepository) GetTrackByMusicianID(musicianID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrackByMusicianID", musicianID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrackByMusicianID indicates an expected call of GetTrackByMusicianID.
func (mr *MockRepositoryMockRecorder) GetTrackByMusicianID(musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrackByMusicianID", reflect.TypeOf((*MockRepository)(nil).GetTrackByMusicianID), musicianID)
}

// GetTracksByAlbumID mocks base method.
func (m *MockRepository) GetTracksByAlbumID(albumID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByAlbumID", albumID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByAlbumID indicates an expected call of GetTracksByAlbumID.
func (mr *MockRepositoryMockRecorder) GetTracksByAlbumID(albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByAlbumID", reflect.TypeOf((*MockRepository)(nil).GetTracksByAlbumID), albumID)
}

// GetTracksByGenreID mocks base method.
func (m *MockRepository) GetTracksByGenreID(genreID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByGenreID", genreID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByGenreID indicates an expected call of GetTracksByGenreID.
func (mr *MockRepositoryMockRecorder) GetTracksByGenreID(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByGenreID", reflect.TypeOf((*MockRepository)(nil).GetTracksByGenreID), genreID)
}

// GetTracksByTittle mocks base method.
func (m *MockRepository) GetTracksByTittle(trackTittle string) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByTittle", trackTittle)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByTittle indicates an expected call of GetTracksByTittle.
func (mr *MockRepositoryMockRecorder) GetTracksByTittle(trackTittle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByTittle", reflect.TypeOf((*MockRepository)(nil).GetTracksByTittle), trackTittle)
}

// GetTracksByUserID mocks base method.
func (m *MockRepository) GetTracksByUserID(userID int) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByUserID", userID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByUserID indicates an expected call of GetTracksByUserID.
func (mr *MockRepositoryMockRecorder) GetTracksByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByUserID", reflect.TypeOf((*MockRepository)(nil).GetTracksByUserID), userID)
}

// SearchTracks mocks base method.
func (m *MockRepository) SearchTracks(searchQuery string) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTracks", searchQuery)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTracks indicates an expected call of SearchTracks.
func (mr *MockRepositoryMockRecorder) SearchTracks(searchQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTracks", reflect.TypeOf((*MockRepository)(nil).SearchTracks), searchQuery)
}

// UploadAudio mocks base method.
func (m *MockRepository) UploadAudio(trackID int, audioPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAudio", trackID, audioPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAudio indicates an expected call of UploadAudio.
func (mr *MockRepositoryMockRecorder) UploadAudio(trackID, audioPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAudio", reflect.TypeOf((*MockRepository)(nil).UploadAudio), trackID, audioPath)
}

// UploadPicture mocks base method.
func (m *MockRepository) UploadPicture(trackID int, audioPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPicture", trackID, audioPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadPicture indicates an expected call of UploadPicture.
func (mr *MockRepositoryMockRecorder) UploadPicture(trackID, audioPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPicture", reflect.TypeOf((*MockRepository)(nil).UploadPicture), trackID, audioPath)
}
