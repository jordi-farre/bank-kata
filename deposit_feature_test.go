package bankkata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func (repository *AccountRepositoryInMemory) GetDeposits() []Deposit {
	return nil
}

func Test_a_deposit(t *testing.T) {
	var repository = AccountRepositoryInMemory{}
	var service = BankService{}
	var deposit = Deposit{Amount: 12.22}
	assert := assert.New(t)

	service.Deposit(deposit)

	var deposits = repository.GetDeposits()
	assert.Contains(deposits, deposit)
}
