// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_musicians is a generated GoMock package.
package mock_musicians

import (
	models "2021_1_Noskool_team/internal/app/musicians/models"
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

// GetMusicianByID mocks base method.
func (m *MockRepository) GetMusicianByID(musicianID int) (*models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusicianByID", musicianID)
	ret0, _ := ret[0].(*models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusicianByID indicates an expected call of GetMusicianByID.
func (mr *MockRepositoryMockRecorder) GetMusicianByID(musicianID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusicianByID", reflect.TypeOf((*MockRepository)(nil).GetMusicianByID), musicianID)
}

// GetMusiciansByGenres mocks base method.
func (m *MockRepository) GetMusiciansByGenres(genre string) (*[]models.Musician, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMusiciansByGenres", genre)
	ret0, _ := ret[0].(*[]models.Musician)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMusiciansByGenres indicates an expected call of GetMusiciansByGenres.
func (mr *MockRepositoryMockRecorder) GetMusiciansByGenres(genre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMusiciansByGenres", reflect.TypeOf((*MockRepository)(nil).GetMusiciansByGenres), genre)
}
