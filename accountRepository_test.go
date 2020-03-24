package bankkata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func (repository *AccountRepositoryInMemory) getDeposits() []Deposit {
	return repository.deposits
}

func Test_save_deposit(t *testing.T) {
	var repository = AccountRepositoryInMemory{}
	var deposit = Deposit{Amount: 12.22}

	repository.Save(deposit)

	var deposits = repository.getDeposits()
	assert.Contains(t, deposits, deposit)
}
