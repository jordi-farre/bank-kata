package bankkata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (repository *AccountRepositoryMock) Save(transaction Transaction) {
	repository.Called(transaction)
}

type ClockMock struct {
	mock.Mock
}

func (clock *ClockMock) Now() time.Time {
	args := clock.Called()
	return args.Get(0).(time.Time)
}

func Test_call_repository_to_save_a_deposit(t *testing.T) {
	var repository AccountRepositoryMock
	var clock ClockMock
	var service = BankService{Repository: &repository, Clock: &clock}
	var deposit = Deposit{Amount: 12.22}
	var date = time.Now()
	var transaction = Transaction{Amount: 12.22, Date: date}
	repository.On("Save", transaction).Return(nil)
	clock.On("Now").Return(date)

	service.Deposit(deposit)

	repository.AssertCalled(t, "Save", transaction)
}

func Test_call_repository_to_save_a_withdrawal(t *testing.T) {
	var repository AccountRepositoryMock
	var clock ClockMock
	var service = BankService{Repository: &repository, Clock: &clock}
	var withdrawal = Withdrawal{Amount: 30.43}
	var date = time.Now()
	var transaction = Transaction{Amount: -30.43, Date: date}
	repository.On("Save", transaction).Return(nil)
	clock.On("Now").Return(date)

	service.Withdrawal(withdrawal)

	repository.AssertCalled(t, "Save", transaction)
}
