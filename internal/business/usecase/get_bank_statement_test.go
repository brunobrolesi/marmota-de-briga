package usecase_test

import (
	"context"
	"testing"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	mock_gateway "github.com/brunobrolesi/marmota-de-briga/mocks/internal_/business/gateway"
	"github.com/stretchr/testify/assert"
)

func TestGetBankStatementUseCase(t *testing.T) {
	type TestSuite struct {
		ClientRepository      *mock_gateway.MockClientRepository
		TransactionRepository *mock_gateway.MockTransactionRepository
		Sut                   usecase.GetBankStatementUseCase
	}

	setup := func(t *testing.T) *TestSuite {
		clientRepository := mock_gateway.NewMockClientRepository(t)
		transactionRepository := mock_gateway.NewMockTransactionRepository(t)
		sut := usecase.NewGetBankStatementUseCase(clientRepository, transactionRepository)

		return &TestSuite{
			ClientRepository:      clientRepository,
			TransactionRepository: transactionRepository,
			Sut:                   sut,
		}
	}

	makeInput := func() *usecase.InputGetBankStatement {
		return &usecase.InputGetBankStatement{ClientID: 1}
	}

	makeTransactions := func() []model.Transaction {
		return []model.Transaction{
			{ClientID: 1, Value: 100, Type: model.Debit, Description: "any_description"},
			{ClientID: 1, Value: 200, Type: model.Credit, Description: "any_description"},
			{ClientID: 1, Value: 300, Type: model.Debit, Description: "any_description"},
			{ClientID: 1, Value: 400, Type: model.Credit, Description: "any_description"},
		}

	}
	t.Run("should return client not found error if client not exits", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		testSuite.ClientRepository.On("GetClient", context.Background(), model.ClientID(1)).Return(nil, model.ErrClientNotFound).Once()
		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientNotFound.Error())
	})
	t.Run("should return internal server error if get client fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		testSuite.ClientRepository.On("GetClient", context.Background(), model.ClientID(1)).Return(nil, model.ErrInternalServerError).Once()
		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})
	t.Run("should return client not found error if client not exists", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		testSuite.ClientRepository.On("GetClient", context.Background(), model.ClientID(1)).Return(nil, nil).Once()
		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientNotFound.Error())
	})
	t.Run("should return internal server error if get last transactions fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), model.ClientID(1)).Return(client, nil).Once()
		testSuite.TransactionRepository.On("GetLastTransactions", context.Background(), model.ClientID(1), model.TRANSACTIONS_LIMIT).Return(nil, model.ErrInternalServerError).Once()
		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		testSuite.TransactionRepository.AssertCalled(t, "GetLastTransactions", context.Background(), input.ClientID, model.TRANSACTIONS_LIMIT)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})
	t.Run("should return bank statement if all goes well", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), model.ClientID(1)).Return(client, nil).Once()
		transactions := makeTransactions()
		testSuite.TransactionRepository.On("GetLastTransactions", context.Background(), model.ClientID(1), model.TRANSACTIONS_LIMIT).Return(transactions, nil).Once()
		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		testSuite.TransactionRepository.AssertCalled(t, "GetLastTransactions", context.Background(), input.ClientID, model.TRANSACTIONS_LIMIT)

		expected := &model.BankStatement{
			Balance: model.BankStatementBalance{
				Total:     model.MonetaryValue(client.AccountBalance),
				CreatedAt: got.Balance.CreatedAt,
				Limit:     model.MonetaryValue(client.AccountLimit),
			},
			Transactions: model.ToBankStatementTransactions(transactions),
		}
		assert.Equal(t, expected, got)
		assert.Nil(t, err)
	})
}
