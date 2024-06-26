// Code generated by mockery v2.40.1. DO NOT EDIT.

package oscar

import (
	context "context"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockAuthService is an autogenerated mock type for the AuthService type
type mockAuthService struct {
	mock.Mock
}

type mockAuthService_Expecter struct {
	mock *mock.Mock
}

func (_m *mockAuthService) EXPECT() *mockAuthService_Expecter {
	return &mockAuthService_Expecter{mock: &_m.Mock}
}

// BUCPChallenge provides a mock function with given fields: bodyIn, newUUID
func (_m *mockAuthService) BUCPChallenge(bodyIn wire.SNAC_0x17_0x06_BUCPChallengeRequest, newUUID func() uuid.UUID) (wire.SNACMessage, error) {
	ret := _m.Called(bodyIn, newUUID)

	if len(ret) == 0 {
		panic("no return value specified for BUCPChallenge")
	}

	var r0 wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(wire.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) (wire.SNACMessage, error)); ok {
		return rf(bodyIn, newUUID)
	}
	if rf, ok := ret.Get(0).(func(wire.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) wire.SNACMessage); ok {
		r0 = rf(bodyIn, newUUID)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(wire.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) error); ok {
		r1 = rf(bodyIn, newUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthService_BUCPChallenge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BUCPChallenge'
type mockAuthService_BUCPChallenge_Call struct {
	*mock.Call
}

// BUCPChallenge is a helper method to define mock.On call
//   - bodyIn wire.SNAC_0x17_0x06_BUCPChallengeRequest
//   - newUUID func() uuid.UUID
func (_e *mockAuthService_Expecter) BUCPChallenge(bodyIn interface{}, newUUID interface{}) *mockAuthService_BUCPChallenge_Call {
	return &mockAuthService_BUCPChallenge_Call{Call: _e.mock.On("BUCPChallenge", bodyIn, newUUID)}
}

func (_c *mockAuthService_BUCPChallenge_Call) Run(run func(bodyIn wire.SNAC_0x17_0x06_BUCPChallengeRequest, newUUID func() uuid.UUID)) *mockAuthService_BUCPChallenge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(wire.SNAC_0x17_0x06_BUCPChallengeRequest), args[1].(func() uuid.UUID))
	})
	return _c
}

func (_c *mockAuthService_BUCPChallenge_Call) Return(_a0 wire.SNACMessage, _a1 error) *mockAuthService_BUCPChallenge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthService_BUCPChallenge_Call) RunAndReturn(run func(wire.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) (wire.SNACMessage, error)) *mockAuthService_BUCPChallenge_Call {
	_c.Call.Return(run)
	return _c
}

// BUCPLogin provides a mock function with given fields: bodyIn, newUUID, fn
func (_m *mockAuthService) BUCPLogin(bodyIn wire.SNAC_0x17_0x02_BUCPLoginRequest, newUUID func() uuid.UUID, fn func(string) (state.User, error)) (wire.SNACMessage, error) {
	ret := _m.Called(bodyIn, newUUID, fn)

	if len(ret) == 0 {
		panic("no return value specified for BUCPLogin")
	}

	var r0 wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(wire.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) (wire.SNACMessage, error)); ok {
		return rf(bodyIn, newUUID, fn)
	}
	if rf, ok := ret.Get(0).(func(wire.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) wire.SNACMessage); ok {
		r0 = rf(bodyIn, newUUID, fn)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(wire.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) error); ok {
		r1 = rf(bodyIn, newUUID, fn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthService_BUCPLogin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BUCPLogin'
type mockAuthService_BUCPLogin_Call struct {
	*mock.Call
}

// BUCPLogin is a helper method to define mock.On call
//   - bodyIn wire.SNAC_0x17_0x02_BUCPLoginRequest
//   - newUUID func() uuid.UUID
//   - fn func(string)(state.User , error)
func (_e *mockAuthService_Expecter) BUCPLogin(bodyIn interface{}, newUUID interface{}, fn interface{}) *mockAuthService_BUCPLogin_Call {
	return &mockAuthService_BUCPLogin_Call{Call: _e.mock.On("BUCPLogin", bodyIn, newUUID, fn)}
}

func (_c *mockAuthService_BUCPLogin_Call) Run(run func(bodyIn wire.SNAC_0x17_0x02_BUCPLoginRequest, newUUID func() uuid.UUID, fn func(string) (state.User, error))) *mockAuthService_BUCPLogin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(wire.SNAC_0x17_0x02_BUCPLoginRequest), args[1].(func() uuid.UUID), args[2].(func(string) (state.User, error)))
	})
	return _c
}

func (_c *mockAuthService_BUCPLogin_Call) Return(_a0 wire.SNACMessage, _a1 error) *mockAuthService_BUCPLogin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthService_BUCPLogin_Call) RunAndReturn(run func(wire.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) (wire.SNACMessage, error)) *mockAuthService_BUCPLogin_Call {
	_c.Call.Return(run)
	return _c
}

// FLAPLogin provides a mock function with given fields: frame, newUUIDFn, newUserFn
func (_m *mockAuthService) FLAPLogin(frame wire.FLAPSignonFrame, newUUIDFn func() uuid.UUID, newUserFn func(string) (state.User, error)) (wire.TLVRestBlock, error) {
	ret := _m.Called(frame, newUUIDFn, newUserFn)

	if len(ret) == 0 {
		panic("no return value specified for FLAPLogin")
	}

	var r0 wire.TLVRestBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(wire.FLAPSignonFrame, func() uuid.UUID, func(string) (state.User, error)) (wire.TLVRestBlock, error)); ok {
		return rf(frame, newUUIDFn, newUserFn)
	}
	if rf, ok := ret.Get(0).(func(wire.FLAPSignonFrame, func() uuid.UUID, func(string) (state.User, error)) wire.TLVRestBlock); ok {
		r0 = rf(frame, newUUIDFn, newUserFn)
	} else {
		r0 = ret.Get(0).(wire.TLVRestBlock)
	}

	if rf, ok := ret.Get(1).(func(wire.FLAPSignonFrame, func() uuid.UUID, func(string) (state.User, error)) error); ok {
		r1 = rf(frame, newUUIDFn, newUserFn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthService_FLAPLogin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FLAPLogin'
type mockAuthService_FLAPLogin_Call struct {
	*mock.Call
}

// FLAPLogin is a helper method to define mock.On call
//   - frame wire.FLAPSignonFrame
//   - newUUIDFn func() uuid.UUID
//   - newUserFn func(string)(state.User , error)
func (_e *mockAuthService_Expecter) FLAPLogin(frame interface{}, newUUIDFn interface{}, newUserFn interface{}) *mockAuthService_FLAPLogin_Call {
	return &mockAuthService_FLAPLogin_Call{Call: _e.mock.On("FLAPLogin", frame, newUUIDFn, newUserFn)}
}

func (_c *mockAuthService_FLAPLogin_Call) Run(run func(frame wire.FLAPSignonFrame, newUUIDFn func() uuid.UUID, newUserFn func(string) (state.User, error))) *mockAuthService_FLAPLogin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(wire.FLAPSignonFrame), args[1].(func() uuid.UUID), args[2].(func(string) (state.User, error)))
	})
	return _c
}

func (_c *mockAuthService_FLAPLogin_Call) Return(_a0 wire.TLVRestBlock, _a1 error) *mockAuthService_FLAPLogin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthService_FLAPLogin_Call) RunAndReturn(run func(wire.FLAPSignonFrame, func() uuid.UUID, func(string) (state.User, error)) (wire.TLVRestBlock, error)) *mockAuthService_FLAPLogin_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveBOSSession provides a mock function with given fields: sessionID
func (_m *mockAuthService) RetrieveBOSSession(sessionID string) (*state.Session, error) {
	ret := _m.Called(sessionID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBOSSession")
	}

	var r0 *state.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*state.Session, error)); ok {
		return rf(sessionID)
	}
	if rf, ok := ret.Get(0).(func(string) *state.Session); ok {
		r0 = rf(sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthService_RetrieveBOSSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveBOSSession'
type mockAuthService_RetrieveBOSSession_Call struct {
	*mock.Call
}

// RetrieveBOSSession is a helper method to define mock.On call
//   - sessionID string
func (_e *mockAuthService_Expecter) RetrieveBOSSession(sessionID interface{}) *mockAuthService_RetrieveBOSSession_Call {
	return &mockAuthService_RetrieveBOSSession_Call{Call: _e.mock.On("RetrieveBOSSession", sessionID)}
}

func (_c *mockAuthService_RetrieveBOSSession_Call) Run(run func(sessionID string)) *mockAuthService_RetrieveBOSSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *mockAuthService_RetrieveBOSSession_Call) Return(_a0 *state.Session, _a1 error) *mockAuthService_RetrieveBOSSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthService_RetrieveBOSSession_Call) RunAndReturn(run func(string) (*state.Session, error)) *mockAuthService_RetrieveBOSSession_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveChatSession provides a mock function with given fields: loginCookie
func (_m *mockAuthService) RetrieveChatSession(loginCookie []byte) (*state.Session, error) {
	ret := _m.Called(loginCookie)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveChatSession")
	}

	var r0 *state.Session
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (*state.Session, error)); ok {
		return rf(loginCookie)
	}
	if rf, ok := ret.Get(0).(func([]byte) *state.Session); ok {
		r0 = rf(loginCookie)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Session)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(loginCookie)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthService_RetrieveChatSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveChatSession'
type mockAuthService_RetrieveChatSession_Call struct {
	*mock.Call
}

// RetrieveChatSession is a helper method to define mock.On call
//   - loginCookie []byte
func (_e *mockAuthService_Expecter) RetrieveChatSession(loginCookie interface{}) *mockAuthService_RetrieveChatSession_Call {
	return &mockAuthService_RetrieveChatSession_Call{Call: _e.mock.On("RetrieveChatSession", loginCookie)}
}

func (_c *mockAuthService_RetrieveChatSession_Call) Run(run func(loginCookie []byte)) *mockAuthService_RetrieveChatSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *mockAuthService_RetrieveChatSession_Call) Return(_a0 *state.Session, _a1 error) *mockAuthService_RetrieveChatSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthService_RetrieveChatSession_Call) RunAndReturn(run func([]byte) (*state.Session, error)) *mockAuthService_RetrieveChatSession_Call {
	_c.Call.Return(run)
	return _c
}

// Signout provides a mock function with given fields: ctx, sess
func (_m *mockAuthService) Signout(ctx context.Context, sess *state.Session) error {
	ret := _m.Called(ctx, sess)

	if len(ret) == 0 {
		panic("no return value specified for Signout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session) error); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAuthService_Signout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Signout'
type mockAuthService_Signout_Call struct {
	*mock.Call
}

// Signout is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
func (_e *mockAuthService_Expecter) Signout(ctx interface{}, sess interface{}) *mockAuthService_Signout_Call {
	return &mockAuthService_Signout_Call{Call: _e.mock.On("Signout", ctx, sess)}
}

func (_c *mockAuthService_Signout_Call) Run(run func(ctx context.Context, sess *state.Session)) *mockAuthService_Signout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session))
	})
	return _c
}

func (_c *mockAuthService_Signout_Call) Return(_a0 error) *mockAuthService_Signout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAuthService_Signout_Call) RunAndReturn(run func(context.Context, *state.Session) error) *mockAuthService_Signout_Call {
	_c.Call.Return(run)
	return _c
}

// SignoutChat provides a mock function with given fields: ctx, sess
func (_m *mockAuthService) SignoutChat(ctx context.Context, sess *state.Session) error {
	ret := _m.Called(ctx, sess)

	if len(ret) == 0 {
		panic("no return value specified for SignoutChat")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session) error); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAuthService_SignoutChat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignoutChat'
type mockAuthService_SignoutChat_Call struct {
	*mock.Call
}

// SignoutChat is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
func (_e *mockAuthService_Expecter) SignoutChat(ctx interface{}, sess interface{}) *mockAuthService_SignoutChat_Call {
	return &mockAuthService_SignoutChat_Call{Call: _e.mock.On("SignoutChat", ctx, sess)}
}

func (_c *mockAuthService_SignoutChat_Call) Run(run func(ctx context.Context, sess *state.Session)) *mockAuthService_SignoutChat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session))
	})
	return _c
}

func (_c *mockAuthService_SignoutChat_Call) Return(_a0 error) *mockAuthService_SignoutChat_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAuthService_SignoutChat_Call) RunAndReturn(run func(context.Context, *state.Session) error) *mockAuthService_SignoutChat_Call {
	_c.Call.Return(run)
	return _c
}

// newMockAuthService creates a new instance of mockAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockAuthService {
	mock := &mockAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
