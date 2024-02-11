package db

import (
	"context"
	"log"

	"github.com/brunobrolesi/marmota-de-briga/cmd/db/cql"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository/config"
	"github.com/brunobrolesi/marmota-de-briga/models"
	"github.com/scylladb/gocqlx/v2/migrate"
)

func InitScyllaDb() {
	createKeyspace()
	migrateKeyspace()
	loadClients()
}

func createKeyspace() {
	client := config.NewScyllaClient("")
	defer client.Close()
	if err := client.ExecStmt(CREATE_KEYSPACE); err != nil {
		log.Fatalln("create keyspace fails with: ", err)
	}
}

func migrateKeyspace() {
	client := config.NewScyllaClient(config.KEYSPACE)
	defer client.Close()
	if err := migrate.FromFS(context.Background(), *client, cql.Files); err != nil {
		log.Fatalln("migrate execution fails with: ", err)
	}
}

func loadClients() {
	client := config.NewScyllaClient(config.KEYSPACE)
	defer client.Close()
	limits := []model.MonetaryValue{100000, 80000, 1000000, 10000000, 500000}
	for i, l := range limits {
		c := model.Client{
			ID:             (i + 1),
			AccountLimit:   l,
			AccountBalance: 0,
		}
		q := client.Query(models.Clients.Insert()).BindStruct(c)
		if err := q.ExecRelease(); err != nil {
			log.Fatalln("load clients fails with: ", err)
		}
	}
}
