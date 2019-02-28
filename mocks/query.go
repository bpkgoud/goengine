// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hellofresh/goengine (interfaces: Query)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	goengine "github.com/hellofresh/goengine"
	reflect "reflect"
)

// Query is a mock of Query interface
type Query struct {
	ctrl     *gomock.Controller
	recorder *QueryMockRecorder
}

// QueryMockRecorder is the mock recorder for Query
type QueryMockRecorder struct {
	mock *Query
}

// NewQuery creates a new mock instance
func NewQuery(ctrl *gomock.Controller) *Query {
	mock := &Query{ctrl: ctrl}
	mock.recorder = &QueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Query) EXPECT() *QueryMockRecorder {
	return m.recorder
}

// Handlers mocks base method
func (m *Query) Handlers() map[string]goengine.MessageHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handlers")
	ret0, _ := ret[0].(map[string]goengine.MessageHandler)
	return ret0
}

// Handlers indicates an expected call of Handlers
func (mr *QueryMockRecorder) Handlers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handlers", reflect.TypeOf((*Query)(nil).Handlers))
}

// Init mocks base method
func (m *Query) Init(arg0 context.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (mr *QueryMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*Query)(nil).Init), arg0)
}
