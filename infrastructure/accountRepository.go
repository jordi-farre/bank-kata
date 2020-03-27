package infrastructure

import "bankkata/domain"

type AccountRepositoryInMemory struct {
	transactions []domain.Transaction
}

func (repository *AccountRepositoryInMemory) Save(transaction domain.Transaction) {
	repository.transactions = append(repository.transactions, transaction)
}

func (repository *AccountRepositoryInMemory) GetTransactions() []domain.Transaction {
	return repository.transactions
}
