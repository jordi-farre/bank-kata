package test

import (
	"bankkata/domain"
	"bankkata/infrastructure"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_save_transaction(t *testing.T) {
	var repository = infrastructure.AccountRepositoryInMemory{}
	var transaction = domain.Transaction{Amount: 12.22, Date: time.Now()}

	repository.Save(transaction)

	var transactions = repository.GetTransactions()
	assert.Contains(t, transactions, transaction)
}
