package usecase

import (
	"context"
	"log"

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
	client, err := uc.clientRepository.GetClient(ctx, input.ClientID)
	if err != nil {
		log.Println("error getting client", err)
		return nil, model.ErrInternalServerError
	}

	if client.AccountBalance.CanNotReceiveTransaction(input.Value, client.AccountLimit, input.Type) {
		log.Println("client limit exceeded")
		return nil, model.ErrClientLimitExceeded
	}

	client.AccountBalance.AddTransaction(input.Value, input.Type)

	err = uc.clientRepository.UpdateBalance(ctx, input.ClientID, client.AccountBalance)
	if err != nil {
		log.Println("error updating client balance", err)
		return nil, model.ErrInternalServerError
	}

	_, err = uc.transactionRepository.CreateTransaction(ctx, input.ClientID, input.Value, input.Type, input.Description)
	if err != nil {
		log.Println("error creating transaction", err)
		return nil, model.ErrInternalServerError
	}

	return client, nil
}
