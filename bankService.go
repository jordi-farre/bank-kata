package bankkata

import (
	"fmt"
	"sort"
	"time"
)

// BankService type
type BankService struct {
	Repository AccountRepository
	Clock      Clock
}

// Deposit blabla
func (bankService *BankService) Deposit(deposit Deposit) {
	bankService.Repository.Save(Transaction{Amount: deposit.Amount, Date: bankService.Clock.Now()})
}

// Withdrawal blabla
func (bankService *BankService) Withdrawal(withdrawal Withdrawal) {
	bankService.Repository.Save(Transaction{Amount: -withdrawal.Amount, Date: bankService.Clock.Now()})
}

func (bankService *BankService) Report() string {
	var transactions = bankService.Repository.GetTransactions()
	var total Amount
	var reportTransactions []ReportTransaction
	for _, transaction := range transactions {
		total += transaction.Amount
		reportTransactions = append(reportTransactions, ReportTransaction{Transaction: transaction, Total: total})
	}
	sort.Sort(sort.Reverse(ByDate(reportTransactions)))
	var report = "date || transaction || balance\n"
	for _, reportTransaction := range reportTransactions {
		report = report + fmt.Sprintf("%s || %v || %v\n", reportTransaction.Date.Format("2006-01-02"), reportTransaction.Amount, reportTransaction.Total)
	}
	return report
}

type ByDate []ReportTransaction

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date.Before(a[j].Date) }

// Amount blabla
type Amount float32

// Deposit blablas
type Deposit struct {
	Amount
}

// Withdrawal blabla
type Withdrawal struct {
	Amount
}

// Transaction blabla
type Transaction struct {
	Amount
	Date time.Time
}

type ReportTransaction struct {
	Transaction
	Total Amount
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
