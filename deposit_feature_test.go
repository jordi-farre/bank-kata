package bankkata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_a_deposit(t *testing.T) {
	var repository = AccountRepositoryInMemory{}
	var service = BankService{Repository: &repository}
	var deposit = Deposit{Amount: 12.22}
	assert := assert.New(t)

	service.Deposit(deposit)

	var deposits = repository.getDeposits()
	assert.Contains(deposits, deposit)
}
