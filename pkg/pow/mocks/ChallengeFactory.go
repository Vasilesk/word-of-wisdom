// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	pow "github.com/vasilesk/word-of-wisdom/pkg/pow"
)

// ChallengeFactory is an autogenerated mock type for the ChallengeFactory type
type ChallengeFactory struct {
	mock.Mock
}

type ChallengeFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *ChallengeFactory) EXPECT() *ChallengeFactory_Expecter {
	return &ChallengeFactory_Expecter{mock: &_m.Mock}
}

// GetNewChallenge provides a mock function with given fields: ctx
func (_m *ChallengeFactory) GetNewChallenge(ctx context.Context) (pow.Challenge, error) {
	ret := _m.Called(ctx)

	var r0 pow.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pow.Challenge, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pow.Challenge); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pow.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChallengeFactory_GetNewChallenge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNewChallenge'
type ChallengeFactory_GetNewChallenge_Call struct {
	*mock.Call
}

// GetNewChallenge is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ChallengeFactory_Expecter) GetNewChallenge(ctx interface{}) *ChallengeFactory_GetNewChallenge_Call {
	return &ChallengeFactory_GetNewChallenge_Call{Call: _e.mock.On("GetNewChallenge", ctx)}
}

func (_c *ChallengeFactory_GetNewChallenge_Call) Run(run func(ctx context.Context)) *ChallengeFactory_GetNewChallenge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ChallengeFactory_GetNewChallenge_Call) Return(_a0 pow.Challenge, _a1 error) *ChallengeFactory_GetNewChallenge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChallengeFactory_GetNewChallenge_Call) RunAndReturn(run func(context.Context) (pow.Challenge, error)) *ChallengeFactory_GetNewChallenge_Call {
	_c.Call.Return(run)
	return _c
}

// RestoreChallenge provides a mock function with given fields: _a0, marshaled
func (_m *ChallengeFactory) RestoreChallenge(_a0 context.Context, marshaled string) (pow.Challenge, error) {
	ret := _m.Called(_a0, marshaled)

	var r0 pow.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (pow.Challenge, error)); ok {
		return rf(_a0, marshaled)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) pow.Challenge); ok {
		r0 = rf(_a0, marshaled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pow.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, marshaled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChallengeFactory_RestoreChallenge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestoreChallenge'
type ChallengeFactory_RestoreChallenge_Call struct {
	*mock.Call
}

// RestoreChallenge is a helper method to define mock.On call
//   - _a0 context.Context
//   - marshaled string
func (_e *ChallengeFactory_Expecter) RestoreChallenge(_a0 interface{}, marshaled interface{}) *ChallengeFactory_RestoreChallenge_Call {
	return &ChallengeFactory_RestoreChallenge_Call{Call: _e.mock.On("RestoreChallenge", _a0, marshaled)}
}

func (_c *ChallengeFactory_RestoreChallenge_Call) Run(run func(_a0 context.Context, marshaled string)) *ChallengeFactory_RestoreChallenge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChallengeFactory_RestoreChallenge_Call) Return(_a0 pow.Challenge, _a1 error) *ChallengeFactory_RestoreChallenge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChallengeFactory_RestoreChallenge_Call) RunAndReturn(run func(context.Context, string) (pow.Challenge, error)) *ChallengeFactory_RestoreChallenge_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewChallengeFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewChallengeFactory creates a new instance of ChallengeFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChallengeFactory(t mockConstructorTestingTNewChallengeFactory) *ChallengeFactory {
	mock := &ChallengeFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
