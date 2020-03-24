package bankkata

// AccountRepository blabla
type AccountRepository interface {
	Save(deposit Deposit)
}

// AccountRepositoryInMemory blabla
type AccountRepositoryInMemory struct {
}
