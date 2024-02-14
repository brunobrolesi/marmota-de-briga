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

		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(&model.Client{}, model.ErrClientNotFound).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientNotFound.Error())
	})

	t.Run("should return internal server error if get client fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(&model.Client{}, model.ErrInternalServerError).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})

	t.Run("should return client limit exceed if transaction exceeds client limit", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 10, AccountBalance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientLimitExceeded.Error())
	})

	t.Run("should return client limit exceed if transaction exceeds client limit", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 10, AccountBalance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientLimitExceeded.Error())
	})

	t.Run("should return internal server error if update balance fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()
		testSuite.ClientRepository.On("UpdateBalance", context.Background(), client, -100).Return(model.ErrInternalServerError).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		testSuite.ClientRepository.AssertCalled(t, "UpdateBalance", context.Background(), client, -100)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})

	t.Run("should return internal server error if create transaction fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()
		testSuite.ClientRepository.On("UpdateBalance", context.Background(), client, -100).Return(nil).Once()
		testSuite.TransactionRepository.On("CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description).Return(nil, model.ErrInternalServerError).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		testSuite.ClientRepository.AssertCalled(t, "UpdateBalance", context.Background(), client, -100)
		testSuite.TransactionRepository.AssertCalled(t, "CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})

	t.Run("should return client if all goes well", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: 0}
		transaction := &model.Transaction{ClientID: input.ClientID, Value: input.Value, Type: input.Type, Description: input.Description}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()
		updatedClient := &model.Client{ID: input.ClientID, AccountLimit: 1000, AccountBalance: -100}
		testSuite.ClientRepository.On("UpdateBalance", context.Background(), client, -100).Return(nil).Once()
		testSuite.TransactionRepository.On("CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description).Return(transaction, nil).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		testSuite.ClientRepository.AssertCalled(t, "UpdateBalance", context.Background(), client, -100)
		testSuite.TransactionRepository.AssertCalled(t, "CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description)
		assert.Nil(t, err)
		assert.Equal(t, updatedClient, got)
	})
}
