package bankkata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_a_deposit(t *testing.T) {
	var repository = AccountRepositoryInMemory{}
	var clock ClockMock
	var service = BankService{Repository: &repository, Clock: &clock}
	var deposit = Deposit{Amount: 12.22}
	var date = time.Now()
	var transaction = Transaction{Amount: 12.22, Date: date}
	assert := assert.New(t)
	clock.On("Now").Return(date)

	service.Deposit(deposit)

	assert.Contains(repository.GetTransactions(), transaction)
}
