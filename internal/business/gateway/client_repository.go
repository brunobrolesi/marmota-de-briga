package gateway

import (
	"context"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
)

type ClientRepository interface {
	GetClient(ctx context.Context, clientID model.ClientID) (*model.Client, error)
	ACIDUpdateBalance(ctx context.Context, clientID model.ClientID, transactionValue model.MonetaryValue, transactionType model.TransactionType) (*model.Client, error)
}
