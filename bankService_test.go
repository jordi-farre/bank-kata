package bankkata

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (repository *AccountRepositoryMock) Save(deposit Deposit) {
	repository.Called(deposit)
}

func Test_call_repository_to_save_a_deposit(t *testing.T) {
	var repository AccountRepositoryMock
	var service = BankService{Repository: &repository}
	var deposit = Deposit{Amount: 12.22}
	repository.On("Save", deposit).Return(nil)

	service.Deposit(deposit)

	repository.AssertCalled(t, "Save", deposit)
}
