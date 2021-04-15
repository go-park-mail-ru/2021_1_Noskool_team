// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_profiles is a generated GoMock package.
package mock_profiles

import (
	models "2021_1_Noskool_team/internal/app/profiles/models"
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

// Create mocks base method.
func (m *MockRepository) Create(u *models.UserProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), u)
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(id string) (*models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*models.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), id)
}

// FindByLogin mocks base method.
func (m *MockRepository) FindByLogin(nickname string) (*models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", nickname)
	ret0, _ := ret[0].(*models.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *MockRepositoryMockRecorder) FindByLogin(nickname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockRepository)(nil).FindByLogin), nickname)
}

// Update mocks base method.
func (m *MockRepository) Update(u *models.UserProfile, withPassword bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", u, withPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(u, withPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), u, withPassword)
}

// UpdateAvatar mocks base method.
func (m *MockRepository) UpdateAvatar(userID, newAvatar string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateAvatar", userID, newAvatar)
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockRepositoryMockRecorder) UpdateAvatar(userID, newAvatar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockRepository)(nil).UpdateAvatar), userID, newAvatar)
}