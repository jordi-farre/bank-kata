package bankkata

// BankService type
type BankService struct {
	Repository AccountRepository
}

// Deposit blabla
func (bankService *BankService) Deposit(deposit Deposit) {
	bankService.Repository.Save(deposit)
}

// Deposit blablas
type Deposit struct {
	Amount float32
}
