package gateway

import (
	"context"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
)

type ClientRepository interface {
	GetClient(ctx context.Context, clientID model.ClientID) (*model.Client, error)
	UpdateBalance(ctx context.Context, client *model.Client, newBalance model.MonetaryValue) error
}
