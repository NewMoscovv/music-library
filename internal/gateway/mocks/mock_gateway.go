// Code generated by MockGen. DO NOT EDIT.
// Source: song_gateway.go
//
// Generated by this command:
//
//	mockgen -source=song_gateway.go -destination=mocks/mock_gateway.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	models "Music-library/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSongGateway is a mock of SongGateway interface.
type MockSongGateway struct {
	ctrl     *gomock.Controller
	recorder *MockSongGatewayMockRecorder
	isgomock struct{}
}

// MockSongGatewayMockRecorder is the mock recorder for MockSongGateway.
type MockSongGatewayMockRecorder struct {
	mock *MockSongGateway
}

// NewMockSongGateway creates a new mock instance.
func NewMockSongGateway(ctrl *gomock.Controller) *MockSongGateway {
	mock := &MockSongGateway{ctrl: ctrl}
	mock.recorder = &MockSongGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSongGateway) EXPECT() *MockSongGatewayMockRecorder {
	return m.recorder
}

// CreateSong mocks base method.
func (m *MockSongGateway) CreateSong(song *models.Song) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSong", song)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSong indicates an expected call of CreateSong.
func (mr *MockSongGatewayMockRecorder) CreateSong(song any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSong", reflect.TypeOf((*MockSongGateway)(nil).CreateSong), song)
}

// DeleteSong mocks base method.
func (m *MockSongGateway) DeleteSong(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSong", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSong indicates an expected call of DeleteSong.
func (mr *MockSongGatewayMockRecorder) DeleteSong(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSong", reflect.TypeOf((*MockSongGateway)(nil).DeleteSong), id)
}

// GetSongByID mocks base method.
func (m *MockSongGateway) GetSongByID(id uint) (*models.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSongByID", id)
	ret0, _ := ret[0].(*models.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSongByID indicates an expected call of GetSongByID.
func (mr *MockSongGatewayMockRecorder) GetSongByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSongByID", reflect.TypeOf((*MockSongGateway)(nil).GetSongByID), id)
}

// GetSongs mocks base method.
func (m *MockSongGateway) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSongs", filter, limit, offset)
	ret0, _ := ret[0].([]models.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSongs indicates an expected call of GetSongs.
func (mr *MockSongGatewayMockRecorder) GetSongs(filter, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSongs", reflect.TypeOf((*MockSongGateway)(nil).GetSongs), filter, limit, offset)
}

// UpdateSong mocks base method.
func (m *MockSongGateway) UpdateSong(song *models.Song) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSong", song)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSong indicates an expected call of UpdateSong.
func (mr *MockSongGatewayMockRecorder) UpdateSong(song any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSong", reflect.TypeOf((*MockSongGateway)(nil).UpdateSong), song)
}
