// Code generated by mockery v2.38.0. DO NOT EDIT.

package handler

import mock "github.com/stretchr/testify/mock"

// mockProfileManager is an autogenerated mock type for the ProfileManager type
type mockProfileManager struct {
	mock.Mock
}

type mockProfileManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockProfileManager) EXPECT() *mockProfileManager_Expecter {
	return &mockProfileManager_Expecter{mock: &_m.Mock}
}

// Profile provides a mock function with given fields: screenName
func (_m *mockProfileManager) Profile(screenName string) (string, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for Profile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(screenName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockProfileManager_Profile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Profile'
type mockProfileManager_Profile_Call struct {
	*mock.Call
}

// Profile is a helper method to define mock.On call
//   - screenName string
func (_e *mockProfileManager_Expecter) Profile(screenName interface{}) *mockProfileManager_Profile_Call {
	return &mockProfileManager_Profile_Call{Call: _e.mock.On("Profile", screenName)}
}

func (_c *mockProfileManager_Profile_Call) Run(run func(screenName string)) *mockProfileManager_Profile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *mockProfileManager_Profile_Call) Return(_a0 string, _a1 error) *mockProfileManager_Profile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockProfileManager_Profile_Call) RunAndReturn(run func(string) (string, error)) *mockProfileManager_Profile_Call {
	_c.Call.Return(run)
	return _c
}

// SetProfile provides a mock function with given fields: screenName, body
func (_m *mockProfileManager) SetProfile(screenName string, body string) error {
	ret := _m.Called(screenName, body)

	if len(ret) == 0 {
		panic("no return value specified for SetProfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(screenName, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockProfileManager_SetProfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetProfile'
type mockProfileManager_SetProfile_Call struct {
	*mock.Call
}

// SetProfile is a helper method to define mock.On call
//   - screenName string
//   - body string
func (_e *mockProfileManager_Expecter) SetProfile(screenName interface{}, body interface{}) *mockProfileManager_SetProfile_Call {
	return &mockProfileManager_SetProfile_Call{Call: _e.mock.On("SetProfile", screenName, body)}
}

func (_c *mockProfileManager_SetProfile_Call) Run(run func(screenName string, body string)) *mockProfileManager_SetProfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *mockProfileManager_SetProfile_Call) Return(_a0 error) *mockProfileManager_SetProfile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockProfileManager_SetProfile_Call) RunAndReturn(run func(string, string) error) *mockProfileManager_SetProfile_Call {
	_c.Call.Return(run)
	return _c
}

// newMockProfileManager creates a new instance of mockProfileManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockProfileManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockProfileManager {
	mock := &mockProfileManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
