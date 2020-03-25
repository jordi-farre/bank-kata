package bankkata

import "time"

// BankService type
type BankService struct {
	Repository AccountRepository
	Clock      Clock
}

// Deposit blabla
func (bankService *BankService) Deposit(deposit Deposit) {
	bankService.Repository.Save(Transaction{Amount: deposit.Amount, Date: bankService.Clock.Now()})
}

// Amount blabla
type Amount float32

// Deposit blablas
type Deposit struct {
	Amount
}

// Transaction blabla
type Transaction struct {
	Amount
	Date time.Time
}

// Clock blabla
type Clock interface {
	Now() time.Time
}

// SystemClock blaba
type SystemClock struct{}

// Now blabla
func (clock *SystemClock) Now() time.Time {
	return time.Now()
}
