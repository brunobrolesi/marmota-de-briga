package repository

import (
	"context"
	"time"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/gateway"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2/log"
)

const (
	queryCreateTransaction   = "INSERT INTO transactions (client_id, value, type, description, created_at) VALUES (?, ?, ?, ?, ?)"
	queryGetLastTransactions = "SELECT client_id, value, type, description, created_at FROM transactions WHERE client_id = ? LIMIT ?"
)

type transactionRepository struct {
	session *gocql.Session
}

func NewTransactionRepository(session *gocql.Session) gateway.TransactionRepository {
	return &transactionRepository{
		session,
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
	q := r.session.Query(queryCreateTransaction, t.ClientID, t.Value, t.Type, t.Description, t.CreatedAt)
	if err := q.Exec(); err != nil {
		log.Error("create transaction fails with: ", err)
		return nil, err
	}
	return &t, nil
}

func (r *transactionRepository) GetLastTransactions(ctx context.Context, clientID model.ClientID, limit uint) ([]model.Transaction, error) {
	transactions := []model.Transaction{}

	s := r.session.Query(queryGetLastTransactions, clientID, limit).Iter().Scanner()

	for s.Next() {
		var transaction model.Transaction
		s.Scan(&transaction.ClientID, &transaction.Value, &transaction.Type, &transaction.Description, &transaction.CreatedAt)
		transactions = append(transactions, transaction)
	}

	if s.Err() != nil {
		log.Error("failed to get last transactions: %v", s.Err())
		return nil, s.Err()
	}

	return transactions, nil
}
