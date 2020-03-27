package test

import (
	"bankkata/domain"
	"bankkata/infrastructure"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_a_deposit(t *testing.T) {
	var repository = infrastructure.AccountRepositoryInMemory{}
	var clock ClockMock
	var service = domain.BankService{Repository: &repository, Clock: &clock}
	var deposit = domain.Deposit{Amount: 12.22}
	var date = time.Now()
	var transaction = domain.Transaction{Amount: 12.22, Date: date}
	assert := assert.New(t)
	clock.On("Now").Return(date)

	service.Deposit(deposit)

	assert.Contains(repository.GetTransactions(), transaction)
}
