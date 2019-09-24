// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	db "github.com/Optum/Redbox/pkg/db"
	mock "github.com/stretchr/testify/mock"
)

// DBer is an autogenerated mock type for the DBer type
type DBer struct {
	mock.Mock
}

// DeleteAccount provides a mock function with given fields: accountID
func (_m *DBer) DeleteAccount(accountID string) (*db.RedboxAccount, error) {
	ret := _m.Called(accountID)

	var r0 *db.RedboxAccount
	if rf, ok := ret.Get(0).(func(string) *db.RedboxAccount); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByPrincipalID provides a mock function with given fields: principalID
func (_m *DBer) FindAccountsByPrincipalID(principalID string) ([]*db.RedboxAccount, error) {
	ret := _m.Called(principalID)

	var r0 []*db.RedboxAccount
	if rf, ok := ret.Get(0).(func(string) []*db.RedboxAccount); ok {
		r0 = rf(principalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(principalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByStatus provides a mock function with given fields: status
func (_m *DBer) FindAccountsByStatus(status db.AccountStatus) ([]*db.RedboxAccount, error) {
	ret := _m.Called(status)

	var r0 []*db.RedboxAccount
	if rf, ok := ret.Get(0).(func(db.AccountStatus) []*db.RedboxAccount); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.AccountStatus) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLeasesByAccount provides a mock function with given fields: accountID
func (_m *DBer) FindLeasesByAccount(accountID string) ([]*db.RedboxLease, error) {
	ret := _m.Called(accountID)

	var r0 []*db.RedboxLease
	if rf, ok := ret.Get(0).(func(string) []*db.RedboxLease); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.RedboxLease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLeasesByPrincipal provides a mock function with given fields: principalID
func (_m *DBer) FindLeasesByPrincipal(principalID string) ([]*db.RedboxLease, error) {
	ret := _m.Called(principalID)

	var r0 []*db.RedboxLease
	if rf, ok := ret.Get(0).(func(string) []*db.RedboxLease); ok {
		r0 = rf(principalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.RedboxLease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(principalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLeasesByStatus provides a mock function with given fields: status
func (_m *DBer) FindLeasesByStatus(status db.LeaseStatus) ([]*db.RedboxLease, error) {
	ret := _m.Called(status)

	var r0 []*db.RedboxLease
	if rf, ok := ret.Get(0).(func(db.LeaseStatus) []*db.RedboxLease); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.RedboxLease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.LeaseStatus) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: accountID
func (_m *DBer) GetAccount(accountID string) (*db.RedboxAccount, error) {
	ret := _m.Called(accountID)

	var r0 *db.RedboxAccount
	if rf, ok := ret.Get(0).(func(string) *db.RedboxAccount); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccounts provides a mock function with given fields:
func (_m *DBer) GetAccounts() ([]*db.RedboxAccount, error) {
	ret := _m.Called()

	var r0 []*db.RedboxAccount
	if rf, ok := ret.Get(0).(func() []*db.RedboxAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountsForReset provides a mock function with given fields:
func (_m *DBer) GetAccountsForReset() ([]*db.RedboxAccount, error) {
	ret := _m.Called()

	var r0 []*db.RedboxAccount
	if rf, ok := ret.Get(0).(func() []*db.RedboxAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReadyAccount provides a mock function with given fields:
func (_m *DBer) GetReadyAccount() (*db.RedboxAccount, error) {
	ret := _m.Called()

	var r0 *db.RedboxAccount
	if rf, ok := ret.Get(0).(func() *db.RedboxAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutAccount provides a mock function with given fields: account
func (_m *DBer) PutAccount(account db.RedboxAccount) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.RedboxAccount) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutLease provides a mock function with given fields: account
func (_m *DBer) PutLease(account db.RedboxLease) (*db.RedboxLease, error) {
	ret := _m.Called(account)

	var r0 *db.RedboxLease
	if rf, ok := ret.Get(0).(func(db.RedboxLease) *db.RedboxLease); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.RedboxLease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.RedboxLease) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransitionAccountStatus provides a mock function with given fields: accountID, prevStatus, nextStatus
func (_m *DBer) TransitionAccountStatus(accountID string, prevStatus db.AccountStatus, nextStatus db.AccountStatus) (*db.RedboxAccount, error) {
	ret := _m.Called(accountID, prevStatus, nextStatus)

	var r0 *db.RedboxAccount
	if rf, ok := ret.Get(0).(func(string, db.AccountStatus, db.AccountStatus) *db.RedboxAccount); ok {
		r0 = rf(accountID, prevStatus, nextStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, db.AccountStatus, db.AccountStatus) error); ok {
		r1 = rf(accountID, prevStatus, nextStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransitionLeaseStatus provides a mock function with given fields: accountID, principalID, prevStatus, nextStatus
func (_m *DBer) TransitionLeaseStatus(accountID string, principalID string, prevStatus db.LeaseStatus, nextStatus db.LeaseStatus) (*db.RedboxLease, error) {
	ret := _m.Called(accountID, principalID, prevStatus, nextStatus)

	var r0 *db.RedboxLease
	if rf, ok := ret.Get(0).(func(string, string, db.LeaseStatus, db.LeaseStatus) *db.RedboxLease); ok {
		r0 = rf(accountID, principalID, prevStatus, nextStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.RedboxLease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, db.LeaseStatus, db.LeaseStatus) error); ok {
		r1 = rf(accountID, principalID, prevStatus, nextStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountPrincipalPolicyHash provides a mock function with given fields: accountID, prevHash, nextHash
func (_m *DBer) UpdateAccountPrincipalPolicyHash(accountID string, prevHash string, nextHash string) (*db.RedboxAccount, error) {
	ret := _m.Called(accountID, prevHash, nextHash)

	var r0 *db.RedboxAccount
	if rf, ok := ret.Get(0).(func(string, string, string) *db.RedboxAccount); ok {
		r0 = rf(accountID, prevHash, nextHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.RedboxAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(accountID, prevHash, nextHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMetadata provides a mock function with given fields: accountID, metadata
func (_m *DBer) UpdateMetadata(accountID string, metadata map[string]interface{}) error {
	ret := _m.Called(accountID, metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) error); ok {
		r0 = rf(accountID, metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}