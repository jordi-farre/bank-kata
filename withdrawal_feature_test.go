package bankkata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_a_withdrawal(t *testing.T) {
	var repository = AccountRepositoryInMemory{}
	var clock ClockMock
	var service = BankService{Repository: &repository, Clock: &clock}
	var withdrawal = Withdrawal{Amount: 25.26}
	var date = time.Now()
	var transaction = Transaction{Amount: -25.26, Date: date}
	assert := assert.New(t)
	clock.On("Now").Return(date)

	service.Withdrawal(withdrawal)

	assert.Contains(repository.GetTransactions(), transaction)
}
