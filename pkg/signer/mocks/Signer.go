// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	signer "github.com/vasilesk/word-of-wisdom/pkg/signer"
)

// Signer is an autogenerated mock type for the Signer type
type Signer struct {
	mock.Mock
}

// Restore provides a mock function with given fields: signed
func (_m *Signer) Restore(signed signer.Signed) (signer.Data, error) {
	ret := _m.Called(signed)

	var r0 signer.Data
	var r1 error
	if rf, ok := ret.Get(0).(func(signer.Signed) (signer.Data, error)); ok {
		return rf(signed)
	}
	if rf, ok := ret.Get(0).(func(signer.Signed) signer.Data); ok {
		r0 = rf(signed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(signer.Data)
		}
	}

	if rf, ok := ret.Get(1).(func(signer.Signed) error); ok {
		r1 = rf(signed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sign provides a mock function with given fields: data
func (_m *Signer) Sign(data signer.Data) (signer.Signed, error) {
	ret := _m.Called(data)

	var r0 signer.Signed
	var r1 error
	if rf, ok := ret.Get(0).(func(signer.Data) (signer.Signed, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(signer.Data) signer.Signed); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(signer.Signed)
		}
	}

	if rf, ok := ret.Get(1).(func(signer.Data) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSigner interface {
	mock.TestingT
	Cleanup(func())
}

// NewSigner creates a new instance of Signer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSigner(t mockConstructorTestingTNewSigner) *Signer {
	mock := &Signer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}