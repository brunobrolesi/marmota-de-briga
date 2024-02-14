package usecase

import (
	"context"
	"errors"

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
	Execute(ctx context.Context, input *InputCreateTransaction) (*model.Client, error)
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

func (uc *createTransactionUseCase) Execute(ctx context.Context, input *InputCreateTransaction) (*model.Client, error) {
	client, err := uc.clientRepository.ACIDUpdateBalance(ctx, input.ClientID, input.Value, input.Type)
	if errors.Is(err, model.ErrClientNotFound) {
		return nil, err
	}

	if errors.Is(err, model.ErrClientLimitExceeded) {
		return nil, err
	}

	if err != nil {
		return nil, model.ErrInternalServerError
	}

	_, err = uc.transactionRepository.CreateTransaction(ctx, input.ClientID, input.Value, input.Type, input.Description)
	if err != nil {
		return nil, model.ErrInternalServerError
	}

	return client, nil
}
