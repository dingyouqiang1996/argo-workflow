// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	jws "github.com/argoproj/argo/server/auth/jws"

	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Authorize provides a mock function with given fields: ctx, authorization
func (_m *Interface) Authorize(ctx context.Context, authorization string) (*jws.ClaimSet, error) {
	ret := _m.Called(ctx, authorization)

	var r0 *jws.ClaimSet
	if rf, ok := ret.Get(0).(func(context.Context, string) *jws.ClaimSet); ok {
		r0 = rf(ctx, authorization)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jws.ClaimSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, authorization)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleCallback provides a mock function with given fields: writer, request
func (_m *Interface) HandleCallback(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// HandleRedirect provides a mock function with given fields: writer, request
func (_m *Interface) HandleRedirect(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// IsRBACEnabled provides a mock function with given fields:
func (_m *Interface) IsRBACEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
