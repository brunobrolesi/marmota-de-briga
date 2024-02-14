package usecase_test

import (
	"context"
	"testing"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	mock_gateway "github.com/brunobrolesi/marmota-de-briga/mocks/internal_/business/gateway"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	type TestSuite struct {
		ClientRepository      *mock_gateway.MockClientRepository
		TransactionRepository *mock_gateway.MockTransactionRepository
		Sut                   usecase.CreateTransactionUseCase
	}

	setup := func(t *testing.T) *TestSuite {
		clientRepository := mock_gateway.NewMockClientRepository(t)
		transactionRepository := mock_gateway.NewMockTransactionRepository(t)
		sut := usecase.NewCreateTransactionUseCase(clientRepository, transactionRepository)

		return &TestSuite{
			ClientRepository:      clientRepository,
			TransactionRepository: transactionRepository,
			Sut:                   sut,
		}
	}

	makeInput := func() *usecase.InputCreateTransaction {
		return &usecase.InputCreateTransaction{
			ClientID:    1,
			Value:       100,
			Type:        model.Debit,
			Description: "any_description",
		}
	}

	t.Run("should return client not found error if client not exists", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		testSuite.ClientRepository.On("ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type).Return(nil, model.ErrClientNotFound).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientNotFound.Error())
	})

	t.Run("should return client limit exceed error if transaction exceeds client limit", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		testSuite.ClientRepository.On("ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type).Return(nil, model.ErrClientLimitExceeded).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientLimitExceeded.Error())
	})

	t.Run("should return internal server error if ACIDUpdateBalance fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		testSuite.ClientRepository.On("ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type).Return(nil, model.ErrInternalServerError).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})

	t.Run("should return internal server error if CreateTransaction fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: -100}
		testSuite.ClientRepository.On("ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type).Return(client, nil).Once()
		testSuite.TransactionRepository.On("CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description).Return(nil, model.ErrInternalServerError).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type)
		testSuite.TransactionRepository.AssertCalled(t, "CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})

	t.Run("should return client if all goes well", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: 0}
		transaction := &model.Transaction{ClientID: input.ClientID, Value: input.Value, Type: input.Type, Description: input.Description}
		testSuite.ClientRepository.On("ACIDUpdateBalance", context.Background(), input.ClientID, input.Value, input.Type).Return(client, nil).Once()
		testSuite.TransactionRepository.On("CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description).Return(transaction, nil).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "ACIDUpdateBalance", context.Background(), client.ID, input.Value, input.Type)
		testSuite.TransactionRepository.AssertCalled(t, "CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description)
		assert.Nil(t, err)
		assert.Equal(t, client, got)
	})
}
