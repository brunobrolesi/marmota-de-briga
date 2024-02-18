package repository

import (
	"context"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/gateway"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2/log"
)

const (
	queryGetClient     = "SELECT id, account_balance, account_limit FROM clients WHERE id = ?"
	queryUpdateBalance = "UPDATE clients SET account_balance = ? WHERE id = ?"
)

type Row = map[string]interface{}

type clientRepository struct {
	session            *gocql.Session
	queryUpdateBalance *gocql.Query
}

func NewClientRepository(session *gocql.Session) gateway.ClientRepository {
	return &clientRepository{
		session:            session,
		queryUpdateBalance: session.Query(queryUpdateBalance),
	}
}

func (r *clientRepository) GetClient(ctx context.Context, id model.ClientID) (*model.Client, error) {
	q := r.session.Query(queryGetClient, id)
	var c model.Client
	if err := q.Scan(&c.ID, &c.AccountBalance, &c.AccountLimit); err != nil {
		if err == gocql.ErrNotFound {
			return nil, model.ErrClientNotFound
		}
		log.Error("get client fails with: ", err)
		return nil, err
	}
	return &c, nil
}

func (r *clientRepository) ACIDUpdateBalance(ctx context.Context, clientID model.ClientID, transactionValue model.MonetaryValue, transactionType model.TransactionType) (*model.Client, error) {
	c, err := r.GetClient(ctx, clientID)
	if err != nil {
		return nil, err
	}
	newBalance, err := c.GetBalanceAfterTransaction(transactionValue, transactionType)
	if err != nil {
		return nil, err
	}
	q := r.queryUpdateBalance.Bind(newBalance, clientID).WithContext(ctx)
	if err := q.Exec(); err != nil {
		log.Error("update balance fails with:", err)
		return nil, err
	}

	c.AccountBalance = newBalance
	return c, nil
}
