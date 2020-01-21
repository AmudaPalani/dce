// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/Optum/dce/pkg/model"

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// GetLeaseByID provides a mock function with given fields: leaseID
func (_m *Reader) GetLeaseByID(leaseID string) (*model.Lease, error) {
	ret := _m.Called(leaseID)

	var r0 *model.Lease
	if rf, ok := ret.Get(0).(func(string) *model.Lease); ok {
		r0 = rf(leaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(leaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeases provides a mock function with given fields: _a0
func (_m *Reader) GetLeases(_a0 *model.Lease) (*model.Leases, error) {
	ret := _m.Called(_a0)

	var r0 *model.Leases
	if rf, ok := ret.Get(0).(func(*model.Lease) *model.Leases); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Leases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Lease) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}