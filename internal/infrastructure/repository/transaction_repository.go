package repository

import (
	"context"
	"log"
	"time"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/gateway"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/models"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
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

func (r *transactionRepository) GetLastTransactions(ctx context.Context, clientID model.ClientID, limit uint) ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	q := qb.Select("transactions").Where(qb.Eq("client_id")).Limit(limit).Query(*r.client).Bind(clientID)
	if err := q.Select(&transactions); err != nil {
		log.Println("get last transactions fails with: ", err)
		return nil, err
	}

	return transactions, nil
}
