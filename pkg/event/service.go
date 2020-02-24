package event

import (
	"github.com/Optum/dce/pkg/account"
	"github.com/Optum/dce/pkg/lease"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// Publisher interface defines anything that can publish an event
type Publisher interface {
	Publish(i interface{}) error
}

// NewServiceInput are the items required to create a new Eventer service
type NewServiceInput struct {
	SnsClient              snsiface.SNSAPI
	SqsClient              sqsiface.SQSAPI
	AccountCreatedTopicArn string `env:"ACCOUNT_CREATED_TOPIC_ARN" envDefault:"arn:aws:sns:us-east-1:123456789012:account-create"`
	AccountDeletedTopicArn string `env:"ACCOUNT_DELETED_TOPIC_ARN" envDefault:"arn:aws:sns:us-east-1:123456789012:account-delete"`
	AccountResetQueueURL   string `env:"RESET_SQS_URL" envDefault:"DefaultResetSQSUrl"`
	LeaseAddedTopicArn     string `env:"LEASE_ADDED_TOPIC" envDefault:"arn:aws:sns:us-east-1:123456789012:lease-added"`
}

// Service is the public interface for publishing events
type Service struct {
	accountCreate []Publisher
	accountDelete []Publisher
	accountUpdate []Publisher
	accountReset  []Publisher
	leaseCreate   []Publisher
	leaseEnd      []Publisher
	leaseUpdate   []Publisher
}

func (e *Service) publish(i interface{}, p ...Publisher) error {
	for _, n := range p {
		err := n.Publish(i)
		if err != nil {
			return err
		}
	}
	return nil
}

// AccountCreate publish events
func (e *Service) AccountCreate(data *account.Account) error {
	return e.publish(data, e.accountCreate...)
}

// AccountDelete publish events
func (e *Service) AccountDelete(data *account.Account) error {
	return e.publish(data, e.accountDelete...)
}

// AccountUpdate publish events
func (e *Service) AccountUpdate(data *account.Account) error {
	return e.publish(data, e.accountUpdate...)
}

// AccountReset publish events
func (e *Service) AccountReset(data *account.Account) error {
	return e.publish(data, e.accountReset...)
}

// LeaseCreate publish events
func (e *Service) LeaseCreate(data *lease.Lease) error {
	return e.publish(data, e.leaseCreate...)
}

// LeaseEnd publish events
func (e *Service) LeaseEnd(i interface{}) error {
	return e.publish(i, e.leaseEnd...)
}

// LeaseUpdate publish events
func (e *Service) LeaseUpdate(i interface{}) error {
	return e.publish(i, e.leaseUpdate...)
}

// NewService creates a new instance of Eventer
func NewService(input NewServiceInput) (*Service, error) {
	newEventer := &Service{}

	//////////////////////////////////////////////////////////////////////
	// Account Eventing
	//////////////////////////////////////////////////////////////////////
	createAccount, err := NewSnsEvent(input.SnsClient, input.AccountCreatedTopicArn)
	if err != nil {
		return nil, err
	}

	resetAccount, err := NewSqsEvent(input.SqsClient, input.AccountResetQueueURL)
	if err != nil {
		return nil, err
	}

	deleteAccount, err := NewSnsEvent(input.SnsClient, input.AccountDeletedTopicArn)
	if err != nil {
		return nil, err
	}
	newEventer.accountCreate = []Publisher{
		createAccount,
	}
	newEventer.accountReset = []Publisher{
		resetAccount,
	}
	newEventer.accountDelete = []Publisher{
		deleteAccount,
	}
	newEventer.accountUpdate = []Publisher{}

	//////////////////////////////////////////////////////////////////////
	// Lease Eventing
	//////////////////////////////////////////////////////////////////////
	createLease, err := NewSnsEvent(input.SnsClient, input.LeaseAddedTopicArn)
	if err != nil {
		return nil, err
	}

	newEventer.leaseCreate = []Publisher{
		createLease,
	}
	newEventer.leaseEnd = []Publisher{}
	newEventer.leaseUpdate = []Publisher{}

	return newEventer, nil
}
