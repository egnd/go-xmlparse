// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	xml "encoding/xml"

	mock "github.com/stretchr/testify/mock"

	xmlparse "github.com/egnd/go-xmlparse"
)

// TokenHandler is an autogenerated mock type for the TokenHandler type
type TokenHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1, _a2
func (_m *TokenHandler) Execute(_a0 interface{}, _a1 xml.StartElement, _a2 xmlparse.TokenReader) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, xml.StartElement, xmlparse.TokenReader) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewTokenHandlerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenHandler creates a new instance of TokenHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenHandler(t NewTokenHandlerT) *TokenHandler {
	mock := &TokenHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
