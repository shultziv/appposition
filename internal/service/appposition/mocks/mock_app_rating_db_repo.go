// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shultziv/appposition/internal/service/appposition (interfaces: AppRatingDbRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/shultziv/appposition/internal/domain"
)

// MockAppRatingDbRepo is a mock of AppRatingDbRepo interface.
type MockAppRatingDbRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAppRatingDbRepoMockRecorder
}

// MockAppRatingDbRepoMockRecorder is the mock recorder for MockAppRatingDbRepo.
type MockAppRatingDbRepoMockRecorder struct {
	mock *MockAppRatingDbRepo
}

// NewMockAppRatingDbRepo creates a new mock instance.
func NewMockAppRatingDbRepo(ctrl *gomock.Controller) *MockAppRatingDbRepo {
	mock := &MockAppRatingDbRepo{ctrl: ctrl}
	mock.recorder = &MockAppRatingDbRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppRatingDbRepo) EXPECT() *MockAppRatingDbRepoMockRecorder {
	return m.recorder
}

// AddAppPositions mocks base method.
func (m *MockAppRatingDbRepo) AddAppPositions(arg0 context.Context, arg1 []*domain.AppPosition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAppPositions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAppPositions indicates an expected call of AddAppPositions.
func (mr *MockAppRatingDbRepoMockRecorder) AddAppPositions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAppPositions", reflect.TypeOf((*MockAppRatingDbRepo)(nil).AddAppPositions), arg0, arg1)
}

// GetMaxPosAppByDays mocks base method.
func (m *MockAppRatingDbRepo) GetMaxPosAppByDays(arg0 context.Context, arg1, arg2 uint32, arg3, arg4 time.Time) (map[uint32]uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxPosAppByDays", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(map[uint32]uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxPosAppByDays indicates an expected call of GetMaxPosAppByDays.
func (mr *MockAppRatingDbRepoMockRecorder) GetMaxPosAppByDays(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxPosAppByDays", reflect.TypeOf((*MockAppRatingDbRepo)(nil).GetMaxPosAppByDays), arg0, arg1, arg2, arg3, arg4)
}