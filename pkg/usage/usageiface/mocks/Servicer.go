// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import usage "github.com/Optum/dce/pkg/usage"

// Servicer is an autogenerated mock type for the Servicer type
type Servicer struct {
	mock.Mock
}

// Get provides a mock function with given fields: startDate, principalID
func (_m *Servicer) Get(startDate int64, principalID string) (*usage.Usage, error) {
	ret := _m.Called(startDate, principalID)

	var r0 *usage.Usage
	if rf, ok := ret.Get(0).(func(int64, string) *usage.Usage); ok {
		r0 = rf(startDate, principalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usage.Usage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(startDate, principalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
