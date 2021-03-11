// Code generated by MockGen. DO NOT EDIT.
// Source: client_interface.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	models "2021_1_Noskool_team/internal/microservices/auth/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthCheckerClient is a mock of AuthCheckerClient interface.
type MockAuthCheckerClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCheckerClientMockRecorder
}

// MockAuthCheckerClientMockRecorder is the mock recorder for MockAuthCheckerClient.
type MockAuthCheckerClientMockRecorder struct {
	mock *MockAuthCheckerClient
}

// NewMockAuthCheckerClient creates a new mock instance.
func NewMockAuthCheckerClient(ctrl *gomock.Controller) *MockAuthCheckerClient {
	mock := &MockAuthCheckerClient{ctrl: ctrl}
	mock.recorder = &MockAuthCheckerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCheckerClient) EXPECT() *MockAuthCheckerClientMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockAuthCheckerClient) Check(ctx context.Context, id string) (models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", ctx, id)
	ret0, _ := ret[0].(models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockAuthCheckerClientMockRecorder) Check(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockAuthCheckerClient)(nil).Check), ctx, id)
}

// Create mocks base method.
func (m *MockAuthCheckerClient) Create(ctx context.Context, id string) (models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, id)
	ret0, _ := ret[0].(models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthCheckerClientMockRecorder) Create(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthCheckerClient)(nil).Create), ctx, id)
}

// Delete mocks base method.
func (m *MockAuthCheckerClient) Delete(ctx context.Context, id string) (models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthCheckerClientMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthCheckerClient)(nil).Delete), ctx, id)
}
