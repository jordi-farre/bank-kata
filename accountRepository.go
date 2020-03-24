package bankkata

// AccountRepository blabla
type AccountRepository interface {
	Save(deposit Deposit)
}

// AccountRepositoryInMemory blabla
type AccountRepositoryInMemory struct {
	deposits []Deposit
}

// Save blabla
func (repository *AccountRepositoryInMemory) Save(deposit Deposit) {
	repository.deposits = append(repository.deposits, deposit)
}
