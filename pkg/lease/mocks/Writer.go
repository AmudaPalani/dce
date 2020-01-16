// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/Optum/dce/pkg/model"

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// WriteLease provides a mock function with given fields: input, lastModifiedOn
func (_m *Writer) WriteLease(input *model.Lease, lastModifiedOn *int64) error {
	ret := _m.Called(input, lastModifiedOn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Lease, *int64) error); ok {
		r0 = rf(input, lastModifiedOn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
