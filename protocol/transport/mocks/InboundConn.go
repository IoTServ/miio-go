// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"
)

// InboundConn is an autogenerated mock type for the InboundConn type
type InboundConn struct {
	mock.Mock
}

// ReadFromUDP provides a mock function with given fields: b
func (_m *InboundConn) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	ret := _m.Called(b)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *net.UDPAddr
	if rf, ok := ret.Get(1).(func([]byte) *net.UDPAddr); ok {
		r1 = rf(b)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*net.UDPAddr)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]byte) error); ok {
		r2 = rf(b)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
