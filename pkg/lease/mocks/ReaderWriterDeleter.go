// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/Optum/dce/pkg/model"

// ReaderWriterDeleter is an autogenerated mock type for the ReaderWriterDeleter type
type ReaderWriterDeleter struct {
	mock.Mock
}

// DeleteLease provides a mock function with given fields: input
func (_m *ReaderWriterDeleter) DeleteLease(input *model.Lease) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Lease) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLeaseByID provides a mock function with given fields: leaseID
func (_m *ReaderWriterDeleter) GetLeaseByID(leaseID string) (*model.Lease, error) {
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
func (_m *ReaderWriterDeleter) GetLeases(_a0 *model.Lease) (*model.Leases, error) {
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

// WriteLease provides a mock function with given fields: input, lastModifiedOn
func (_m *ReaderWriterDeleter) WriteLease(input *model.Lease, lastModifiedOn *int64) error {
	ret := _m.Called(input, lastModifiedOn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Lease, *int64) error); ok {
		r0 = rf(input, lastModifiedOn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}