// Code generated by mockery v2.43.2. DO NOT EDIT.

package handler

import (
	context "context"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockChatService is an autogenerated mock type for the ChatService type
type mockChatService struct {
	mock.Mock
}

type mockChatService_Expecter struct {
	mock *mock.Mock
}

func (_m *mockChatService) EXPECT() *mockChatService_Expecter {
	return &mockChatService_Expecter{mock: &_m.Mock}
}

// ChannelMsgToHost provides a mock function with given fields: ctx, sess, inFrame, inBody
func (_m *mockChatService) ChannelMsgToHost(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*wire.SNACMessage, error) {
	ret := _m.Called(ctx, sess, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for ChannelMsgToHost")
	}

	var r0 *wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*wire.SNACMessage, error)); ok {
		return rf(ctx, sess, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x0E_0x05_ChatChannelMsgToHost) *wire.SNACMessage); ok {
		r0 = rf(ctx, sess, inFrame, inBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wire.SNACMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x0E_0x05_ChatChannelMsgToHost) error); ok {
		r1 = rf(ctx, sess, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockChatService_ChannelMsgToHost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelMsgToHost'
type mockChatService_ChannelMsgToHost_Call struct {
	*mock.Call
}

// ChannelMsgToHost is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inFrame wire.SNACFrame
//   - inBody wire.SNAC_0x0E_0x05_ChatChannelMsgToHost
func (_e *mockChatService_Expecter) ChannelMsgToHost(ctx interface{}, sess interface{}, inFrame interface{}, inBody interface{}) *mockChatService_ChannelMsgToHost_Call {
	return &mockChatService_ChannelMsgToHost_Call{Call: _e.mock.On("ChannelMsgToHost", ctx, sess, inFrame, inBody)}
}

func (_c *mockChatService_ChannelMsgToHost_Call) Run(run func(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x0E_0x05_ChatChannelMsgToHost)) *mockChatService_ChannelMsgToHost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(wire.SNAC_0x0E_0x05_ChatChannelMsgToHost))
	})
	return _c
}

func (_c *mockChatService_ChannelMsgToHost_Call) Return(_a0 *wire.SNACMessage, _a1 error) *mockChatService_ChannelMsgToHost_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockChatService_ChannelMsgToHost_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*wire.SNACMessage, error)) *mockChatService_ChannelMsgToHost_Call {
	_c.Call.Return(run)
	return _c
}

// newMockChatService creates a new instance of mockChatService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockChatService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockChatService {
	mock := &mockChatService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
