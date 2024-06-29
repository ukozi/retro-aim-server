// Code generated by mockery v2.43.2. DO NOT EDIT.

package http

import (
	context "context"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockMessageRelayer is an autogenerated mock type for the MessageRelayer type
type mockMessageRelayer struct {
	mock.Mock
}

type mockMessageRelayer_Expecter struct {
	mock *mock.Mock
}

func (_m *mockMessageRelayer) EXPECT() *mockMessageRelayer_Expecter {
	return &mockMessageRelayer_Expecter{mock: &_m.Mock}
}

// RelayToScreenName provides a mock function with given fields: ctx, screenName, msg
func (_m *mockMessageRelayer) RelayToScreenName(ctx context.Context, screenName state.IdentScreenName, msg wire.SNACMessage) {
	_m.Called(ctx, screenName, msg)
}

// mockMessageRelayer_RelayToScreenName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RelayToScreenName'
type mockMessageRelayer_RelayToScreenName_Call struct {
	*mock.Call
}

// RelayToScreenName is a helper method to define mock.On call
//   - ctx context.Context
//   - screenName state.IdentScreenName
//   - msg wire.SNACMessage
func (_e *mockMessageRelayer_Expecter) RelayToScreenName(ctx interface{}, screenName interface{}, msg interface{}) *mockMessageRelayer_RelayToScreenName_Call {
	return &mockMessageRelayer_RelayToScreenName_Call{Call: _e.mock.On("RelayToScreenName", ctx, screenName, msg)}
}

func (_c *mockMessageRelayer_RelayToScreenName_Call) Run(run func(ctx context.Context, screenName state.IdentScreenName, msg wire.SNACMessage)) *mockMessageRelayer_RelayToScreenName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(state.IdentScreenName), args[2].(wire.SNACMessage))
	})
	return _c
}

func (_c *mockMessageRelayer_RelayToScreenName_Call) Return() *mockMessageRelayer_RelayToScreenName_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockMessageRelayer_RelayToScreenName_Call) RunAndReturn(run func(context.Context, state.IdentScreenName, wire.SNACMessage)) *mockMessageRelayer_RelayToScreenName_Call {
	_c.Call.Return(run)
	return _c
}

// newMockMessageRelayer creates a new instance of mockMessageRelayer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockMessageRelayer(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockMessageRelayer {
	mock := &mockMessageRelayer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
