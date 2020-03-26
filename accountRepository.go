package bankkata

// AccountRepository blabla
type AccountRepository interface {
	Save(transaction Transaction)
	GetTransactions() []Transaction
}

// AccountRepositoryInMemory blabla
type AccountRepositoryInMemory struct {
	transactions []Transaction
}

// Save blabla
func (repository *AccountRepositoryInMemory) Save(transaction Transaction) {
	repository.transactions = append(repository.transactions, transaction)
}

func (repository *AccountRepositoryInMemory) GetTransactions() []Transaction {
	return repository.transactions
}
