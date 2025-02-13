// Code generated by mockery v2.44.1. DO NOT EDIT.

package auth

import (
	context "context"

	auth "github.com/DNS-Pro/core/internal/auth"

	mock "github.com/stretchr/testify/mock"
)

// MockIAuther is an autogenerated mock type for the IAuther type
type MockIAuther struct {
	mock.Mock
}

type MockIAuther_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIAuther) EXPECT() *MockIAuther_Expecter {
	return &MockIAuther_Expecter{mock: &_m.Mock}
}

// GetType provides a mock function with given fields:
func (_m *MockIAuther) GetType() auth.AuthType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetType")
	}

	var r0 auth.AuthType
	if rf, ok := ret.Get(0).(func() auth.AuthType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(auth.AuthType)
	}

	return r0
}

// MockIAuther_GetType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetType'
type MockIAuther_GetType_Call struct {
	*mock.Call
}

// GetType is a helper method to define mock.On call
func (_e *MockIAuther_Expecter) GetType() *MockIAuther_GetType_Call {
	return &MockIAuther_GetType_Call{Call: _e.mock.On("GetType")}
}

func (_c *MockIAuther_GetType_Call) Run(run func()) *MockIAuther_GetType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIAuther_GetType_Call) Return(_a0 auth.AuthType) *MockIAuther_GetType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIAuther_GetType_Call) RunAndReturn(run func() auth.AuthType) *MockIAuther_GetType_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockIAuther) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIAuther_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockIAuther_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockIAuther_Expecter) Run(ctx interface{}) *MockIAuther_Run_Call {
	return &MockIAuther_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockIAuther_Run_Call) Run(run func(ctx context.Context)) *MockIAuther_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIAuther_Run_Call) Return(_a0 error) *MockIAuther_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIAuther_Run_Call) RunAndReturn(run func(context.Context) error) *MockIAuther_Run_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIAuther creates a new instance of MockIAuther. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIAuther(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIAuther {
	mock := &MockIAuther{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
