package repository

import (
	"context"
	"log"
	"time"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/gateway"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/models"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/v2"
)

type transactionRepository struct {
	client *gocqlx.Session
}

func NewTransactionRepository(client *gocqlx.Session) gateway.TransactionRepository {
	return &transactionRepository{
		client: client,
	}
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, clientID model.ClientID, value model.MonetaryValue, transactionType model.TransactionType, description string) (*model.Transaction, error) {
	t := model.Transaction{
		ID:          uuid.New(),
		ClientID:    clientID,
		Value:       value,
		Type:        transactionType,
		Description: description,
		CreatedAt:   time.Now(),
	}
	q := r.client.Query(models.Transactions.Insert()).BindStruct(t)
	if err := q.ExecRelease(); err != nil {
		log.Println("create transaction fails with: ", err)
		return nil, err
	}
	return &t, nil
}

func (r *transactionRepository) GetLastTransactions(ctx context.Context, clientID model.ClientID, limit int) ([]model.Transaction, error) {
	return nil, nil
}
