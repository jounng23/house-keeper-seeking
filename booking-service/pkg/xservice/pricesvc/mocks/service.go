// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package pricesvc_mocks is a generated GoMock package.
package pricesvc_mocks

import (
	pricesvc "booking-svc/pkg/xservice/pricesvc"
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

// GetPrice mocks base method.
func (m *MockService) GetPrice(c context.Context, datetime int64) (pricesvc.GetPriceReponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice", c, datetime)
	ret0, _ := ret[0].(pricesvc.GetPriceReponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockServiceMockRecorder) GetPrice(c, datetime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockService)(nil).GetPrice), c, datetime)
}
