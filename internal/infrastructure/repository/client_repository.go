package repository

import (
	"context"
	"log"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/gateway"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/models"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type clientRepository struct {
	client *gocqlx.Session
}

func NewClientRepository(client *gocqlx.Session) gateway.ClientRepository {
	return &clientRepository{
		client: client,
	}
}

func (r *clientRepository) GetClient(ctx context.Context, id int) (*model.Client, error) {
	c := model.Client{
		ID: id,
	}
	q := r.client.Query(models.Clients.Get()).BindStruct(c)
	if err := q.GetRelease(&c); err != nil {
		log.Println("get client fails with: ", err)
		return nil, err
	}
	return &c, nil
}

func (r *clientRepository) UpdateBalance(ctx context.Context, clientID model.ClientID, newBalance model.ClientBalance) error {
	c := model.Client{
		ID:             clientID,
		AccountBalance: newBalance,
	}

	q := qb.Update("clients").Set("account_balance").Where(qb.Eq("id")).Query(*r.client).BindStruct(c)
	if err := q.ExecRelease(); err != nil {
		log.Println("update client balance fails with: ", err)
		return err
	}
	return nil
}
