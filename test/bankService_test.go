package test

import (
	"bankkata/domain"
	"bankkata/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (repository *AccountRepositoryMock) Save(transaction domain.Transaction) {
	repository.Called(transaction)
}

func (repository *AccountRepositoryMock) GetTransactions() []domain.Transaction {
	args := repository.Called()
	return args.Get(0).([]domain.Transaction)
}

func Test_call_repository_to_save_a_deposit(t *testing.T) {
	var repository AccountRepositoryMock
	var clock mocks.ClockMock
	var service = domain.BankService{Repository: &repository, Clock: &clock}
	var deposit = domain.Deposit{Amount: 12.22}
	var date = time.Now()
	var transaction = domain.Transaction{Amount: 12.22, Date: date}
	repository.On("Save", transaction).Return(nil)
	clock.On("Now").Return(date)

	service.Deposit(deposit)

	repository.AssertCalled(t, "Save", transaction)
}

func Test_call_repository_to_save_a_withdrawal(t *testing.T) {
	var repository AccountRepositoryMock
	var clock mocks.ClockMock
	var service = domain.BankService{Repository: &repository, Clock: &clock}
	var withdrawal = domain.Withdrawal{Amount: 30.43}
	var date = time.Now()
	var transaction = domain.Transaction{Amount: -30.43, Date: date}
	repository.On("Save", transaction).Return(nil)
	clock.On("Now").Return(date)

	service.Withdrawal(withdrawal)

	repository.AssertCalled(t, "Save", transaction)
}

func Test_call_repository_to_get_all_transactions(t *testing.T) {
	var repository AccountRepositoryMock
	var service = domain.BankService{Repository: &repository}
	repository.On("GetTransactions").Return([]domain.Transaction{domain.Transaction{Amount: -30.43, Date: time.Now()}})

	service.Report()

	repository.AssertCalled(t, "GetTransactions")
}

func Test_return_a_report_of_transactions(t *testing.T) {
	var repository AccountRepositoryMock
	var service = domain.BankService{Repository: &repository}
	var date, _ = time.Parse("2006-01-02", "2020-02-11")
	repository.On("GetTransactions").Return([]domain.Transaction{
		domain.Transaction{Amount: 5000, Date: date},
		domain.Transaction{Amount: -2000, Date: date.AddDate(0, 0, 4)},
		domain.Transaction{Amount: 10000, Date: date.AddDate(0, 0, 7)},
	})

	var report = service.Report()

	assert.Equal(t, "date || transaction || balance\n2020-02-18 || 10000 || 13000\n2020-02-15 || -2000 || 3000\n2020-02-11 || 5000 || 5000\n", report)
}
