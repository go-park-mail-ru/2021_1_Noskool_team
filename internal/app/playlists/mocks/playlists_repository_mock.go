// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_playlists is a generated GoMock package.
package mock_playlists

import (
	models "2021_1_Noskool_team/internal/app/playlists/models"
	models0 "2021_1_Noskool_team/internal/app/tracks/models"
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

// AddPlaylistToMediateka mocks base method.
func (m *MockRepository) AddPlaylistToMediateka(userID, playlistID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlaylistToMediateka", userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlaylistToMediateka indicates an expected call of AddPlaylistToMediateka.
func (mr *MockRepositoryMockRecorder) AddPlaylistToMediateka(userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlaylistToMediateka", reflect.TypeOf((*MockRepository)(nil).AddPlaylistToMediateka), userID, playlistID)
}

// CreatePlaylist mocks base method.
func (m *MockRepository) CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", playlist)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist indicates an expected call of CreatePlaylist.
func (mr *MockRepositoryMockRecorder) CreatePlaylist(playlist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockRepository)(nil).CreatePlaylist), playlist)
}

// DeletePlaylistFromUser mocks base method.
func (m *MockRepository) DeletePlaylistFromUser(userID, playlistID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylistFromUser", userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlaylistFromUser indicates an expected call of DeletePlaylistFromUser.
func (mr *MockRepositoryMockRecorder) DeletePlaylistFromUser(userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylistFromUser", reflect.TypeOf((*MockRepository)(nil).DeletePlaylistFromUser), userID, playlistID)
}

// GetMediateka mocks base method.
func (m *MockRepository) GetMediateka(userID int) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMediateka", userID)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMediateka indicates an expected call of GetMediateka.
func (mr *MockRepositoryMockRecorder) GetMediateka(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediateka", reflect.TypeOf((*MockRepository)(nil).GetMediateka), userID)
}

// GetPlaylistByID mocks base method.
func (m *MockRepository) GetPlaylistByID(playlistID int) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylistByID", playlistID)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylistByID indicates an expected call of GetPlaylistByID.
func (mr *MockRepositoryMockRecorder) GetPlaylistByID(playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylistByID", reflect.TypeOf((*MockRepository)(nil).GetPlaylistByID), playlistID)
}

// GetPlaylistsByGenreID mocks base method.
func (m *MockRepository) GetPlaylistsByGenreID(genreID int) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylistsByGenreID", genreID)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylistsByGenreID indicates an expected call of GetPlaylistsByGenreID.
func (mr *MockRepositoryMockRecorder) GetPlaylistsByGenreID(genreID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylistsByGenreID", reflect.TypeOf((*MockRepository)(nil).GetPlaylistsByGenreID), genreID)
}

// GetTracksByPlaylistID mocks base method.
func (m *MockRepository) GetTracksByPlaylistID(playlistID int) ([]*models0.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracksByPlaylistID", playlistID)
	ret0, _ := ret[0].([]*models0.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTracksByPlaylistID indicates an expected call of GetTracksByPlaylistID.
func (mr *MockRepositoryMockRecorder) GetTracksByPlaylistID(playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracksByPlaylistID", reflect.TypeOf((*MockRepository)(nil).GetTracksByPlaylistID), playlistID)
}

// SearchPlaylists mocks base method.
func (m *MockRepository) SearchPlaylists(searchQuery string) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPlaylists", searchQuery)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPlaylists indicates an expected call of SearchPlaylists.
func (mr *MockRepositoryMockRecorder) SearchPlaylists(searchQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPlaylists", reflect.TypeOf((*MockRepository)(nil).SearchPlaylists), searchQuery)
}

// UploadPicture mocks base method.
func (m *MockRepository) UploadPicture(playlistID int, audioPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPicture", playlistID, audioPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadPicture indicates an expected call of UploadPicture.
func (mr *MockRepositoryMockRecorder) UploadPicture(playlistID, audioPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPicture", reflect.TypeOf((*MockRepository)(nil).UploadPicture), playlistID, audioPath)
}