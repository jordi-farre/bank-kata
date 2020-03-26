package bankkata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_a_bank_service(t *testing.T) {
	var repository = AccountRepositoryInMemory{}
	var clock ClockMock
	var service = BankService{Repository: &repository, Clock: &clock}
	clock.On("Now").Times(1).Return(time.Parse("2006-01-02", "2020-01-10"))
	service.Deposit(Deposit{Amount: 1000})
	clock.On("Now").Times(1).Return(time.Parse("2006-01-02", "2020-01-13"))
	service.Deposit(Deposit{Amount: 2000})
	clock.On("Now").Times(1).Return(time.Parse("2006-01-02", "2020-01-14"))
	service.Withdrawal(Withdrawal{Amount: 500})

	var report = service.Report()

	assert.Equal(t, "date || transaction || balance\n2020-01-14 || -500 || 2500\n2020-01-13 || 2000 || 3000\n2020-01-10 || 1000 || 1000\n", report)
}
