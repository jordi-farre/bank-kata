package mocks

import (
	"bankkata/domain"

	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (repository *AccountRepositoryMock) Save(transaction domain.Transaction) {
	repository.Called(transaction)
}

func (repository *AccountRepositoryMock) GetTransactions() []domain.Transaction {
	args := repository.Called()
	return args.Get(0).([]domain.Transaction)
}
