package main

import (
	"fmt"
	"bankkata/domain"
	"bankkata/infrastructure"
)

func main() {
	var clock = infrastructure.SystemClock{}
	var repository = infrastructure.AccountRepositoryInMemory{}
	var service = domain.BankService{Repository: &repository, Clock: &clock}
	service.Deposit(domain.Deposit{Amount: 12500})
	service.Withdrawal(domain.Withdrawal{Amount: 10000})
	service.Deposit(domain.Deposit{Amount: 5000})
	service.Withdrawal(domain.Withdrawal{Amount: 3000})
	service.Deposit(domain.Deposit{Amount: 2225})
	fmt.Print(service.Report())
}