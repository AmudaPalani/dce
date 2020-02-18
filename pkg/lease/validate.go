package lease

import (
	"errors"
	"reflect"
	"regexp"

	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"math"
	"time"
)

// We don't use the internal errors package here because validation will rewrite it anyways
// Just spit out errors and turn them into validation errors inside the appropriate functions

var validateID = []validation.Rule{
	validation.NotNil.Error("must be a string"),
	is.UUIDv4.Error("must be a UUIDv4"),
}

var validateAccountID = []validation.Rule{
	validation.NotNil.Error("must be a string"),
	validation.Match(regexp.MustCompile("^[0-9]{12}$")).Error("must be a string with 12 digits"),
}

var validatePrincipalID = []validation.Rule{
	validation.NotNil.Error("must be a string"),
}

var validateInt64 = []validation.Rule{
	validation.NotNil.Error("must be an epoch timestamp"),
}

var validateStatus = []validation.Rule{
	validation.NotNil.Error("must be a valid lease status"),
}

func isNil(value interface{}) error {
	if !reflect.ValueOf(value).IsNil() {
		return errors.New("must be empty")
	}
	return nil
}

func isLeaseActive(value interface{}) error {
	s, _ := value.(*Status)
	if s.String() != StatusActive.String() {
		return errors.New("must be active lease")
	}
	return nil
}

func isExpiresOnValid(a *Service) validation.RuleFunc {

	return func(value interface{}) error {
		if !reflect.ValueOf(value).IsNil() {

			e, _ := value.(*int64)

			// Validate requested lease end date is greater than today
			if *e <= time.Now().Unix() {
				return fmt.Errorf("Requested lease has a desired expiry date less than today: %d", *e)
			}

			// Validate requested lease budget period is less than MAX_LEASE_BUDGET_PERIOD
			currentTime := time.Now()
			maxLeaseExpiresOn := currentTime.Add(time.Second * time.Duration(a.maxLeasePeriod))
			if *e > maxLeaseExpiresOn.Unix() {
				return fmt.Errorf("Requested lease has a budget expires on of %d, which is greater than max lease period of %d", *e, a.maxLeasePeriod)
			}
		}
		return nil
	}
}

func isBudgetAmountValid(a *Service) validation.RuleFunc {
	return func(value interface{}) error {
		if !reflect.ValueOf(value).IsNil() {
			b, _ := value.(*float64)

			// Validate requested lease budget amount is less than MAX_LEASE_BUDGET_AMOUNT
			if *b > a.maxLeaseBudgetAmount {
				return fmt.Errorf("Requested lease has a budget amount of %f, which is greater than max lease budget amount of %f", math.Round(*b), math.Round(a.maxLeaseBudgetAmount))
			}

			/*
				// Validate requested lease budget amount is less than PRINCIPAL_BUDGET_AMOUNT for current principal billing period
				usageStartTime := getBeginningOfCurrentBillingPeriod(a.principalBudgetPeriod)

					usageRecords, err := usageSvc.GetUsageByPrincipal(usageStartTime, *requestBody.PrincipalID)
					if err != nil {
						errStr := fmt.Sprintf("Failed to retrieve usage: %s", err)
						return requestBody, true, "", errors.New(errStr)
					}

					// Group by PrincipalID to get sum of total spent for current billing period
					spent := 0.0
					for _, usageItem := range usageRecords {
						spent = spent + *usageItem.CostAmount
					}

					if spent > a.principalBudgetAmount {
						validationErrStr := fmt.Sprintf(
							"Unable to create lease: User principal %s has already spent %.2f of their %.2f principal budget",
							*requestBody.PrincipalID, spent, a.principalBudgetAmount,
						)
						return requestBody, false, validationErrStr, nil
					}
			*/
		}
		return nil
	}
}

// getBeginningOfCurrentBillingPeriod returns starts of the billing period based on budget period
func getBeginningOfCurrentBillingPeriod(input string) time.Time {
	currentTime := time.Now()
	if input == Weekly {

		for currentTime.Weekday() != time.Sunday { // iterate back to Sunday
			currentTime = currentTime.AddDate(0, 0, -1)
		}

		return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)
	}

	return time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
}
