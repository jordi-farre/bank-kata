package bankkata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func (repository *AccountRepositoryInMemory) getTransactions() []Transaction {
	return repository.transactions
}

func Test_save_deposit(t *testing.T) {
	var repository = AccountRepositoryInMemory{}
	var transaction = Transaction{Amount: 12.22, Date: time.Now()}

	repository.Save(transaction)

	var transactions = repository.getTransactions()
	assert.Contains(t, transactions, transaction)
}
