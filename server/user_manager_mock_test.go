// Code generated by mockery v2.38.0. DO NOT EDIT.

package server

import (
	state "github.com/mkaminski/goaim/state"
	mock "github.com/stretchr/testify/mock"
)

// mockUserManager is an autogenerated mock type for the UserManager type
type mockUserManager struct {
	mock.Mock
}

type mockUserManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockUserManager) EXPECT() *mockUserManager_Expecter {
	return &mockUserManager_Expecter{mock: &_m.Mock}
}

// AllUsers provides a mock function with given fields:
func (_m *mockUserManager) AllUsers() ([]state.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AllUsers")
	}

	var r0 []state.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]state.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []state.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserManager_AllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllUsers'
type mockUserManager_AllUsers_Call struct {
	*mock.Call
}

// AllUsers is a helper method to define mock.On call
func (_e *mockUserManager_Expecter) AllUsers() *mockUserManager_AllUsers_Call {
	return &mockUserManager_AllUsers_Call{Call: _e.mock.On("AllUsers")}
}

func (_c *mockUserManager_AllUsers_Call) Run(run func()) *mockUserManager_AllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockUserManager_AllUsers_Call) Return(_a0 []state.User, _a1 error) *mockUserManager_AllUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserManager_AllUsers_Call) RunAndReturn(run func() ([]state.User, error)) *mockUserManager_AllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// InsertUser provides a mock function with given fields: u
func (_m *mockUserManager) InsertUser(u state.User) error {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for InsertUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.User) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockUserManager_InsertUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertUser'
type mockUserManager_InsertUser_Call struct {
	*mock.Call
}

// InsertUser is a helper method to define mock.On call
//   - u state.User
func (_e *mockUserManager_Expecter) InsertUser(u interface{}) *mockUserManager_InsertUser_Call {
	return &mockUserManager_InsertUser_Call{Call: _e.mock.On("InsertUser", u)}
}

func (_c *mockUserManager_InsertUser_Call) Run(run func(u state.User)) *mockUserManager_InsertUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.User))
	})
	return _c
}

func (_c *mockUserManager_InsertUser_Call) Return(_a0 error) *mockUserManager_InsertUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockUserManager_InsertUser_Call) RunAndReturn(run func(state.User) error) *mockUserManager_InsertUser_Call {
	_c.Call.Return(run)
	return _c
}

// newMockUserManager creates a new instance of mockUserManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockUserManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockUserManager {
	mock := &mockUserManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
