package test

import (
	"bankkata/domain"
	"bankkata/infrastructure"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_a_withdrawal(t *testing.T) {
	var repository = infrastructure.AccountRepositoryInMemory{}
	var clock ClockMock
	var service = domain.BankService{Repository: &repository, Clock: &clock}
	var withdrawal = domain.Withdrawal{Amount: 25.26}
	var date = time.Now()
	var transaction = domain.Transaction{Amount: -25.26, Date: date}
	assert := assert.New(t)
	clock.On("Now").Return(date)

	service.Withdrawal(withdrawal)

	assert.Contains(repository.GetTransactions(), transaction)
}
