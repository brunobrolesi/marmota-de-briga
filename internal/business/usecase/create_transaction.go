package usecase

import (
	"context"
	"fmt"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/gateway"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
)

type InputCreateTransaction struct {
	ClientID    model.ClientID
	Value       model.MonetaryValue
	Type        model.TransactionType
	Description string
}

type CreateTransactionUseCase interface {
	Execute(ctx context.Context, input *InputCreateTransaction) (*model.Transaction, error)
}

type createTransactionUseCase struct {
	clientRepository      gateway.ClientRepository
	transactionRepository gateway.TransactionRepository
}

func NewCreateTransactionUseCase(
	clientRepository gateway.ClientRepository,
	transactionRepository gateway.TransactionRepository,
) CreateTransactionUseCase {
	return &createTransactionUseCase{
		clientRepository:      clientRepository,
		transactionRepository: transactionRepository,
	}
}

func (c *createTransactionUseCase) Execute(ctx context.Context, input *InputCreateTransaction) (*model.Transaction, error) {
	client, err := c.clientRepository.GetClient(ctx, input.ClientID)
	if err != nil {
		fmt.Println("error getting client limit", err)
		return nil, model.ErrInternalServerError
	}

	if client.Balance.CanNotReceiveTransaction(input.Value, client.Limit, input.Type) {
		fmt.Println("client limit exceeded")
		return nil, model.ErrClientLimitExceeded
	}

	transaction, err := c.transactionRepository.CreateTransaction(ctx, input.ClientID, input.Value, input.Type, input.Description)
	if err != nil {
		fmt.Println("error creating transaction", err)
		return nil, model.ErrInternalServerError
	}

	return transaction, nil
}
