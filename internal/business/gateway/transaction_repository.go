package gateway

import (
	"context"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, clientID model.ClientID, value model.MonetaryValue, transactionType model.TransactionType, description string) (*model.Transaction, error)
}
