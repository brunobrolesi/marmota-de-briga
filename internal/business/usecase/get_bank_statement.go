package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/gateway"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
)

type InputGetBankStatement struct {
	ClientID model.ClientID
}

type GetBankStatementUseCase interface {
	Execute(ctx context.Context, input *InputGetBankStatement) (*model.BankStatement, error)
}

type getBankStatementUseCase struct {
	clientRepository      gateway.ClientRepository
	transactionRepository gateway.TransactionRepository
}

func NewGetBankStatementUseCase(
	clientRepository gateway.ClientRepository,
	transactionRepository gateway.TransactionRepository,
) GetBankStatementUseCase {
	return &getBankStatementUseCase{
		clientRepository:      clientRepository,
		transactionRepository: transactionRepository,
	}
}

func (uc *getBankStatementUseCase) Execute(ctx context.Context, input *InputGetBankStatement) (*model.BankStatement, error) {
	client, err := uc.clientRepository.GetClient(ctx, input.ClientID)
	if err != nil {
		fmt.Println("error getting client", err)
		return nil, model.ErrInternalServerError
	}

	if client == nil {
		return nil, model.ErrClientNotFound
	}

	transactions, err := uc.transactionRepository.GetLastTransactions(ctx, input.ClientID, model.TRANSACTIONS_LIMIT)
	if err != nil {
		fmt.Println("error getting transactions", err)
		return nil, model.ErrInternalServerError
	}

	balance := model.BankStatementBalance{
		Total:     model.MonetaryValue(client.Balance),
		CreatedAt: time.Now(),
		Limit:     client.Limit,
	}

	bankStatement := &model.BankStatement{
		Balance:      balance,
		Transactions: model.ToBankStatementTransactions(transactions),
	}

	return bankStatement, nil
}
