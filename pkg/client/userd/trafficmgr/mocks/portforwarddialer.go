// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/telepresenceio/telepresence/v2/pkg/dnet (interfaces: PortForwardDialer)

// Package mock_trafficmgr is a generated GoMock package.
package mock_trafficmgr

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPortForwardDialer is a mock of PortForwardDialer interface.
type MockPortForwardDialer struct {
	ctrl     *gomock.Controller
	recorder *MockPortForwardDialerMockRecorder
}

// MockPortForwardDialerMockRecorder is the mock recorder for MockPortForwardDialer.
type MockPortForwardDialerMockRecorder struct {
	mock *MockPortForwardDialer
}

// NewMockPortForwardDialer creates a new mock instance.
func NewMockPortForwardDialer(ctrl *gomock.Controller) *MockPortForwardDialer {
	mock := &MockPortForwardDialer{ctrl: ctrl}
	mock.recorder = &MockPortForwardDialerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortForwardDialer) EXPECT() *MockPortForwardDialerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPortForwardDialer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPortForwardDialerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPortForwardDialer)(nil).Close))
}

// Dial mocks base method.
func (m *MockPortForwardDialer) Dial(arg0 context.Context, arg1 string) (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dial", arg0, arg1)
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dial indicates an expected call of Dial.
func (mr *MockPortForwardDialerMockRecorder) Dial(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dial", reflect.TypeOf((*MockPortForwardDialer)(nil).Dial), arg0, arg1)
}