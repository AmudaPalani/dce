package lease_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Optum/dce/pkg/errors"
	"github.com/Optum/dce/pkg/lease"
	"github.com/Optum/dce/pkg/lease/mocks"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func ptrString(s string) *string {
	ptrS := s
	return &ptrS
}

func ptrFloat(s float64) *float64 {
	ptrS := s
	return &ptrS
}
func ptrArrayString(s []string) *[]string {
	ptrS := s
	return &ptrS
}

func TestGetLeaseByID(t *testing.T) {

	type response struct {
		data *lease.Lease
		err  error
	}

	tests := []struct {
		name string
		ID   string
		ret  response
		exp  response
	}{
		{
			name: "should get an lease by ID",
			ID:   "70c2d96d-7938-4ec9-917d-476f2b09cc04",
			ret: response{
				data: &lease.Lease{
					ID:     ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
					Status: lease.StatusActive.StatusPtr(),
				},
				err: nil,
			},
			exp: response{
				data: &lease.Lease{
					ID:     ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
					Status: lease.StatusActive.StatusPtr(),
				},
				err: nil,
			},
		},
		{
			name: "should get failure",
			ret: response{
				data: nil,
				err:  errors.NewInternalServer("failure", fmt.Errorf("original failure")),
			},
			exp: response{
				data: nil,
				err:  errors.NewInternalServer("failure", fmt.Errorf("original failure")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksRwd := &mocks.ReaderWriterDeleter{}

			mocksRwd.On("Get", tt.ID).Return(tt.ret.data, tt.ret.err)

			leaseSvc := lease.NewService(lease.NewServiceInput{
				DataSvc: mocksRwd,
			})

			getLease, err := leaseSvc.Get(tt.ID)
			assert.True(t, errors.Is(err, tt.exp.err), "actual error %q doesn't match expected error %q", err, tt.exp.err)

			assert.Equal(t, tt.exp.data, getLease)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		ID        string
		expErr    error
		returnErr error
		expLease  *lease.Lease
	}{
		{
			name: "should delete a lease",
			ID:   "70c2d96d-7938-4ec9-917d-476f2b09cc04",
			expLease: &lease.Lease{
				ID:           ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
				Status:       lease.StatusActive.StatusPtr(),
				StatusReason: lease.StatusReasonDestroyed.StatusReasonPtr(),
			},
			returnErr: nil,
		},
		{
			name:      "should error when delete fails",
			ID:        "70c2d96d-7938-4ec9-917d-476f2b09cc04",
			expLease:  nil,
			returnErr: errors.NewInternalServer("failure", fmt.Errorf("original failure")),
			expErr:    errors.NewInternalServer("failure", nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksRwd := &mocks.ReaderWriterDeleter{}
			mocksRwd.On("Get", tt.ID).
				Return(tt.expLease, tt.returnErr)

			mocksRwd.On("Write", mock.Anything, mock.Anything).
				Return(tt.returnErr)

			leaseSvc := lease.NewService(
				lease.NewServiceInput{
					DataSvc: mocksRwd,
				},
			)
			actualLease, err := leaseSvc.Delete(tt.ID)
			assert.True(t, errors.Is(err, tt.expErr), "actual error %q doesn't match expected error %q", err, tt.expErr)
			assert.Equal(t, tt.expLease, actualLease)

		})
	}
}

func TestSave(t *testing.T) {
	now := time.Now().Unix()

	type response struct {
		data *lease.Lease
		err  error
	}

	tests := []struct {
		name      string
		returnErr error
		lease     *lease.Lease
		exp       response
	}{
		{
			name: "should save lease with timestamps",
			lease: &lease.Lease{
				ID:             ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
				Status:         lease.StatusActive.StatusPtr(),
				AccountID:      ptrString("123456789012"),
				PrincipalID:    ptrString("test:arn"),
				CreatedOn:      &now,
				LastModifiedOn: &now,
			},
			exp: response{
				data: &lease.Lease{
					ID:             ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
					Status:         lease.StatusActive.StatusPtr(),
					AccountID:      ptrString("123456789012"),
					PrincipalID:    ptrString("test:arn"),
					LastModifiedOn: &now,
					CreatedOn:      &now,
				},
				err: nil,
			},
			returnErr: nil,
		},
		{
			name: "should save with new created on",
			lease: &lease.Lease{
				ID:          ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
				Status:      lease.StatusActive.StatusPtr(),
				PrincipalID: ptrString("test:arn"),
				AccountID:   ptrString("123456789012"),
			},
			exp: response{
				data: &lease.Lease{
					ID:             ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
					Status:         lease.StatusActive.StatusPtr(),
					PrincipalID:    ptrString("test:arn"),
					AccountID:      ptrString("123456789012"),
					LastModifiedOn: &now,
					CreatedOn:      &now,
				},
				err: nil,
			},
			returnErr: nil,
		},
		{
			name: "should fail on return err",
			lease: &lease.Lease{
				ID:          ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
				Status:      lease.StatusActive.StatusPtr(),
				PrincipalID: ptrString("test:arn"),
				AccountID:   ptrString("123456789012"),
			},
			exp: response{
				data: &lease.Lease{
					ID:             ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
					Status:         lease.StatusActive.StatusPtr(),
					PrincipalID:    ptrString("test:arn"),
					AccountID:      ptrString("123456789012"),
					LastModifiedOn: &now,
					CreatedOn:      &now,
				},
				err: errors.NewInternalServer("failure", nil),
			},
			returnErr: errors.NewInternalServer("failure", nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksRwd := &mocks.ReaderWriterDeleter{}

			mocksRwd.On("Write", mock.AnythingOfType("*lease.Lease"), mock.AnythingOfType("*int64")).Return(tt.returnErr)

			leaseSvc := lease.NewService(
				lease.NewServiceInput{
					DataSvc: mocksRwd,
				},
			)

			err := leaseSvc.Save(tt.lease)

			assert.Truef(t, errors.Is(err, tt.exp.err), "actual error %q doesn't match expected error %q", err, tt.exp.err)
			assert.Equal(t, tt.exp.data, tt.lease)

		})
	}
}

func TestGetLeases(t *testing.T) {

	type response struct {
		data *lease.Leases
		err  error
	}

	tests := []struct {
		name      string
		inputData lease.Lease
		ret       response
		exp       response
	}{
		{
			name: "standard",
			inputData: lease.Lease{
				Status: lease.StatusActive.StatusPtr(),
			},
			ret: response{
				data: &lease.Leases{
					lease.Lease{
						ID:     aws.String("1"),
						Status: lease.StatusActive.StatusPtr(),
					},
					lease.Lease{
						ID:     aws.String("2"),
						Status: lease.StatusActive.StatusPtr(),
					},
				},
				err: nil,
			},
			exp: response{
				data: &lease.Leases{
					lease.Lease{
						ID:     ptrString("1"),
						Status: lease.StatusActive.StatusPtr(),
					},
					lease.Lease{
						ID:     ptrString("2"),
						Status: lease.StatusActive.StatusPtr(),
					},
				},
				err: nil,
			},
		},
		{
			name: "internal error",
			inputData: lease.Lease{
				Status: lease.StatusActive.StatusPtr(),
			},
			ret: response{
				data: nil,
				err:  errors.NewInternalServer("failure", fmt.Errorf("original error")),
			},
			exp: response{
				data: nil,
				err:  errors.NewInternalServer("failure", fmt.Errorf("original error")),
			},
		},
		{
			name: "validation error",
			inputData: lease.Lease{
				ID: ptrString("70c2d96d-7938-4ec9-917d-476f2b09cc04"),
			},
			ret: response{
				data: nil,
				err:  nil,
			},
			exp: response{
				data: nil,
				err:  errors.NewValidation("lease", fmt.Errorf("id: must be empty.")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocksRWD := &mocks.ReaderWriterDeleter{}
			mocksRWD.On("List", mock.AnythingOfType("*lease.Lease")).Return(tt.ret.data, tt.ret.err)

			leasesSvc := lease.NewService(
				lease.NewServiceInput{
					DataSvc: mocksRWD,
				},
			)

			leases, err := leasesSvc.List(&tt.inputData)
			assert.True(t, errors.Is(err, tt.exp.err), "actual error %q doesn't match expected error %q", err, tt.exp.err)
			assert.Equal(t, tt.exp.data, leases)
		})
	}

}

func TestCreate(t *testing.T) {

	type response struct {
		data *lease.Lease
		err  error
	}

	leaseExpires := time.Now().AddDate(0, 0, 7).Unix()

	tests := []struct {
		name           string
		req            *lease.Lease
		exp            response
		getResponse    *lease.Leases
		writeErr       error
		leaseCreateErr error
	}{
		{
			name: "should create",
			req: &lease.Lease{
				PrincipalID:              ptrString("User1"),
				AccountID:                ptrString("123456789012"),
				BudgetAmount:             ptrFloat(200.00),
				BudgetCurrency:           ptrString("USD"),
				BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
				Metadata:                 map[string]interface{}{},
			},
			exp: response{
				data: &lease.Lease{
					//ID:                       ptrString("6d666a28-4f2c-43af-8c94-1b715ca079ae"),
					PrincipalID:              ptrString("User1"),
					AccountID:                ptrString("123456789012"),
					BudgetAmount:             ptrFloat(200.00),
					BudgetCurrency:           ptrString("USD"),
					BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
					ExpiresOn:                &leaseExpires,
					Metadata:                 map[string]interface{}{},
				},
				err: nil,
			},
			getResponse:    nil,
			writeErr:       nil,
			leaseCreateErr: nil,
		},
		{
			name: "should fail on lease validation error caused by budget amount greater than max lease budget amount",
			req: &lease.Lease{
				PrincipalID:              ptrString("User1"),
				AccountID:                ptrString("123456789012"),
				BudgetAmount:             ptrFloat(2000.00),
				BudgetCurrency:           ptrString("USD"),
				BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
				Metadata:                 map[string]interface{}{},
			},
			exp: response{
				data: nil,
				err:  errors.NewValidation("lease", fmt.Errorf("Requested lease has a budget amount of 2000.000000, which is greater than max lease budget amount of 1000.000000")),
			},
			getResponse: &lease.Leases{
				lease.Lease{
					PrincipalID:              ptrString("User1"),
					AccountID:                ptrString("123456789012"),
					BudgetAmount:             ptrFloat(200.00),
					BudgetCurrency:           ptrString("USD"),
					BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
					Metadata:                 map[string]interface{}{},
				},
			},
		},
		{
			name: "should fail on lease validation error caused by expires on not being nil",
			req: &lease.Lease{
				PrincipalID:              ptrString("User1"),
				AccountID:                ptrString("123456789012"),
				BudgetAmount:             ptrFloat(2000.00),
				BudgetCurrency:           ptrString("USD"),
				BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
				Metadata:                 map[string]interface{}{},
				ExpiresOn:                &leaseExpires,
			},
			exp: response{
				data: nil,
				err:  errors.NewValidation("lease", fmt.Errorf("expiresOn: must be empty.")),
			},
			getResponse: &lease.Leases{
				lease.Lease{
					PrincipalID:              ptrString("User1"),
					AccountID:                ptrString("123456789012"),
					BudgetAmount:             ptrFloat(200.00),
					BudgetCurrency:           ptrString("USD"),
					BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
					Metadata:                 map[string]interface{}{},
				},
			},
		},
		{
			name: "should fail on lease already exists",
			req: &lease.Lease{
				PrincipalID:              ptrString("User1"),
				AccountID:                ptrString("123456789012"),
				BudgetAmount:             ptrFloat(200.00),
				BudgetCurrency:           ptrString("USD"),
				BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
				Metadata:                 map[string]interface{}{},
			},
			exp: response{
				data: nil,
				err:  errors.NewAlreadyExists("lease", "User1"),
			},
			getResponse: &lease.Leases{
				lease.Lease{
					PrincipalID:              ptrString("User1"),
					AccountID:                ptrString("123456789012"),
					Status:                   lease.StatusActive.StatusPtr(),
					BudgetAmount:             ptrFloat(200.00),
					BudgetCurrency:           ptrString("USD"),
					BudgetNotificationEmails: ptrArrayString([]string{"test1@test.com", "test2@test.com"}),
					Metadata:                 map[string]interface{}{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mocksRwd := &mocks.ReaderWriterDeleter{}
			mocksRwd.On("List", mock.AnythingOfType("*lease.Lease")).Return(tt.getResponse, nil)
			mocksRwd.On("Write", mock.AnythingOfType("*lease.Lease"), mock.AnythingOfType("*int64")).Return(tt.writeErr)

			leaseSvc := lease.NewService(
				lease.NewServiceInput{
					DataSvc:                  mocksRwd,
					DefaultLeaseLengthInDays: 7,
					PrincipalBudgetAmount:    1000.00,
					PrincipalBudgetPeriod:    "Weekly",
					MaxLeaseBudgetAmount:     1000.00,
					MaxLeasePeriod:           704800,
				},
			)

			result, err := leaseSvc.Create(tt.req)

			assert.Truef(t, errors.Is(err, tt.exp.err), "actual error %q doesn't match expected error %q", err, tt.exp.err)
			assert.Equal(t, tt.exp.data, result)
		})
	}
}
