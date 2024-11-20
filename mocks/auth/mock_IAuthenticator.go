// Code generated by mockery v2.44.1. DO NOT EDIT.

package auth

import (
	context "context"

	auth "github.com/DNS-Pro/core/internal/auth"

	mock "github.com/stretchr/testify/mock"
)

// MockIAuthenticator is an autogenerated mock type for the IAuthenticator type
type MockIAuthenticator struct {
	mock.Mock
}

type MockIAuthenticator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIAuthenticator) EXPECT() *MockIAuthenticator_Expecter {
	return &MockIAuthenticator_Expecter{mock: &_m.Mock}
}

// Run provides a mock function with given fields: ctx
func (_m *MockIAuthenticator) Run(ctx context.Context) error {
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

// MockIAuthenticator_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockIAuthenticator_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockIAuthenticator_Expecter) Run(ctx interface{}) *MockIAuthenticator_Run_Call {
	return &MockIAuthenticator_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockIAuthenticator_Run_Call) Run(run func(ctx context.Context)) *MockIAuthenticator_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIAuthenticator_Run_Call) Return(_a0 error) *MockIAuthenticator_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIAuthenticator_Run_Call) RunAndReturn(run func(context.Context) error) *MockIAuthenticator_Run_Call {
	_c.Call.Return(run)
	return _c
}

// SetBaseAuth provides a mock function with given fields: _a0
func (_m *MockIAuthenticator) SetBaseAuth(_a0 *auth.Authenticator) {
	_m.Called(_a0)
}

// MockIAuthenticator_SetBaseAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBaseAuth'
type MockIAuthenticator_SetBaseAuth_Call struct {
	*mock.Call
}

// SetBaseAuth is a helper method to define mock.On call
//   - _a0 *auth.Authenticator
func (_e *MockIAuthenticator_Expecter) SetBaseAuth(_a0 interface{}) *MockIAuthenticator_SetBaseAuth_Call {
	return &MockIAuthenticator_SetBaseAuth_Call{Call: _e.mock.On("SetBaseAuth", _a0)}
}

func (_c *MockIAuthenticator_SetBaseAuth_Call) Run(run func(_a0 *auth.Authenticator)) *MockIAuthenticator_SetBaseAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*auth.Authenticator))
	})
	return _c
}

func (_c *MockIAuthenticator_SetBaseAuth_Call) Return() *MockIAuthenticator_SetBaseAuth_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIAuthenticator_SetBaseAuth_Call) RunAndReturn(run func(*auth.Authenticator)) *MockIAuthenticator_SetBaseAuth_Call {
	_c.Call.Return(run)
	return _c
}

// SetDefaults provides a mock function with given fields:
func (_m *MockIAuthenticator) SetDefaults() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SetDefaults")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIAuthenticator_SetDefaults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDefaults'
type MockIAuthenticator_SetDefaults_Call struct {
	*mock.Call
}

// SetDefaults is a helper method to define mock.On call
func (_e *MockIAuthenticator_Expecter) SetDefaults() *MockIAuthenticator_SetDefaults_Call {
	return &MockIAuthenticator_SetDefaults_Call{Call: _e.mock.On("SetDefaults")}
}

func (_c *MockIAuthenticator_SetDefaults_Call) Run(run func()) *MockIAuthenticator_SetDefaults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIAuthenticator_SetDefaults_Call) Return(_a0 error) *MockIAuthenticator_SetDefaults_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIAuthenticator_SetDefaults_Call) RunAndReturn(run func() error) *MockIAuthenticator_SetDefaults_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields:
func (_m *MockIAuthenticator) Validate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIAuthenticator_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type MockIAuthenticator_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
func (_e *MockIAuthenticator_Expecter) Validate() *MockIAuthenticator_Validate_Call {
	return &MockIAuthenticator_Validate_Call{Call: _e.mock.On("Validate")}
}

func (_c *MockIAuthenticator_Validate_Call) Run(run func()) *MockIAuthenticator_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIAuthenticator_Validate_Call) Return(_a0 error) *MockIAuthenticator_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIAuthenticator_Validate_Call) RunAndReturn(run func() error) *MockIAuthenticator_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIAuthenticator creates a new instance of MockIAuthenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIAuthenticator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIAuthenticator {
	mock := &MockIAuthenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}