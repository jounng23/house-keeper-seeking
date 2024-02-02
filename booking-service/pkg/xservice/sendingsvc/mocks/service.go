// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package sendingsvc_mocks is a generated GoMock package.
package sendingsvc_mocks

import (
	sendingsvc "booking-svc/pkg/xservice/sendingsvc"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// PostNotification mocks base method.
func (m *MockService) PostNotification(c context.Context, req sendingsvc.PostNotificationRequest) (sendingsvc.BaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostNotification", c, req)
	ret0, _ := ret[0].(sendingsvc.BaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostNotification indicates an expected call of PostNotification.
func (mr *MockServiceMockRecorder) PostNotification(c, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostNotification", reflect.TypeOf((*MockService)(nil).PostNotification), c, req)
}
