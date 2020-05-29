// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/zitadel/internal/notification/providers (interfaces: NotificationProvider)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNotificationProvider is a mock of NotificationProvider interface
type MockNotificationProvider struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationProviderMockRecorder
}

// MockNotificationProviderMockRecorder is the mock recorder for MockNotificationProvider
type MockNotificationProviderMockRecorder struct {
	mock *MockNotificationProvider
}

// NewMockNotificationProvider creates a new mock instance
func NewMockNotificationProvider(ctrl *gomock.Controller) *MockNotificationProvider {
	mock := &MockNotificationProvider{ctrl: ctrl}
	mock.recorder = &MockNotificationProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotificationProvider) EXPECT() *MockNotificationProviderMockRecorder {
	return m.recorder
}

// CanHandleMessage mocks base method
func (m *MockNotificationProvider) CanHandleMessage() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanHandleMessage")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanHandleMessage indicates an expected call of CanHandleMessage
func (mr *MockNotificationProviderMockRecorder) CanHandleMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanHandleMessage", reflect.TypeOf((*MockNotificationProvider)(nil).CanHandleMessage))
}

// HandleMessage mocks base method
func (m *MockNotificationProvider) HandleMessage() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage")
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage
func (mr *MockNotificationProviderMockRecorder) HandleMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockNotificationProvider)(nil).HandleMessage))
}