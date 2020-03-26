package bankkata

type AccountRepository interface {
	Save(transaction Transaction)
	GetTransactions() []Transaction
}

type AccountRepositoryInMemory struct {
	transactions []Transaction
}

func (repository *AccountRepositoryInMemory) Save(transaction Transaction) {
	repository.transactions = append(repository.transactions, transaction)
}

func (repository *AccountRepositoryInMemory) GetTransactions() []Transaction {
	return repository.transactions
}
