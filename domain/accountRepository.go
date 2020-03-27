package domain

type AccountRepository interface {
	Save(transaction Transaction)
	GetTransactions() []Transaction
}
