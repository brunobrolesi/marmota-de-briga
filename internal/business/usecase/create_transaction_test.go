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

		client := &model.Client{ID: input.ClientID, Limit: 10, Balance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientLimitExceeded.Error())
	})

	t.Run("should return client limit exceed if transaction exceeds client limit", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, Limit: 10, Balance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrClientLimitExceeded.Error())
	})

	t.Run("should return internal server error if create transaction fails", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, Limit: 1000, Balance: 0}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()
		testSuite.TransactionRepository.On("CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description).Return(nil, model.ErrInternalServerError).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		testSuite.TransactionRepository.AssertCalled(t, "CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description)
		assert.Nil(t, got)
		assert.EqualError(t, err, model.ErrInternalServerError.Error())
	})

	t.Run("should return transaction if all goes well", func(t *testing.T) {
		testSuite := setup(t)
		input := makeInput()

		client := &model.Client{ID: input.ClientID, Limit: 1000, Balance: 0}
		transaction := &model.Transaction{ID: 1, ClientID: input.ClientID, Value: input.Value, Type: input.Type, Description: input.Description}
		testSuite.ClientRepository.On("GetClient", context.Background(), input.ClientID).Return(client, nil).Once()
		testSuite.TransactionRepository.On("CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description).Return(transaction, nil).Once()

		got, err := testSuite.Sut.Execute(context.Background(), input)

		testSuite.ClientRepository.AssertCalled(t, "GetClient", context.Background(), input.ClientID)
		testSuite.TransactionRepository.AssertCalled(t, "CreateTransaction", context.Background(), input.ClientID, input.Value, input.Type, input.Description)
		assert.Nil(t, err)
		assert.Equal(t, transaction, got)
	})
}
