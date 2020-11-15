// Code generated by MockGen. DO NOT EDIT.
// Source: ../svc/account/internal/usecase/usecase.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
)

// MockrepoInterface is a mock of repoInterface interface.
type MockrepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockrepoInterfaceMockRecorder
}

// MockrepoInterfaceMockRecorder is the mock recorder for MockrepoInterface.
type MockrepoInterfaceMockRecorder struct {
	mock *MockrepoInterface
}

// NewMockrepoInterface creates a new mock instance.
func NewMockrepoInterface(ctrl *gomock.Controller) *MockrepoInterface {
	mock := &MockrepoInterface{ctrl: ctrl}
	mock.recorder = &MockrepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockrepoInterface) EXPECT() *MockrepoInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockrepoInterface) Get(ctx context.Context, param entity.ParamGetUser) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, param)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockrepoInterfaceMockRecorder) Get(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockrepoInterface)(nil).Get), ctx, param)
}

// Save mocks base method.
func (m *MockrepoInterface) Save(ctx context.Context, data entity.ParamSaveUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockrepoInterfaceMockRecorder) Save(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockrepoInterface)(nil).Save), ctx, data)
}
