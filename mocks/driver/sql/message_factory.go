// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hellofresh/goengine/driver/sql (interfaces: MessageFactory)

// Package sql is a generated GoMock package.
package sql

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	goengine "github.com/hellofresh/goengine"
	reflect "reflect"
)

// MessageFactory is a mock of MessageFactory interface
type MessageFactory struct {
	ctrl     *gomock.Controller
	recorder *MessageFactoryMockRecorder
}

// MessageFactoryMockRecorder is the mock recorder for MessageFactory
type MessageFactoryMockRecorder struct {
	mock *MessageFactory
}

// NewMessageFactory creates a new mock instance
func NewMessageFactory(ctrl *gomock.Controller) *MessageFactory {
	mock := &MessageFactory{ctrl: ctrl}
	mock.recorder = &MessageFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MessageFactory) EXPECT() *MessageFactoryMockRecorder {
	return m.recorder
}

// CreateEventStream mocks base method
func (m *MessageFactory) CreateEventStream(arg0 *sql.Rows) (goengine.EventStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventStream", arg0)
	ret0, _ := ret[0].(goengine.EventStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEventStream indicates an expected call of CreateEventStream
func (mr *MessageFactoryMockRecorder) CreateEventStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventStream", reflect.TypeOf((*MessageFactory)(nil).CreateEventStream), arg0)
}
