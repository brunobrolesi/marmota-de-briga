package db

import (
	"log"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository/config"
)

func InitScyllaDb() {
	createKeyspace()
	migrateKeyspace()
	loadClients()
}

func createKeyspace() {
	client := config.GetScyllaSession("")
	defer client.Close()
	if err := client.Query(CREATE_KEYSPACE).Exec(); err != nil {
		log.Fatalln("create keyspace fails with: ", err)
	}
}

func migrateKeyspace() {
	session := config.GetScyllaSession(config.KEYSPACE)
	defer session.Close()
	if err := session.Query(CREATE_CLIENTS).Exec(); err != nil {
		log.Fatalln("create clients table fails with: ", err)
	}
	if err := session.Query(CREATE_TRANSACTIONS).Exec(); err != nil {
		log.Fatalln("create transactions table fails with: ", err)
	}
	if err := session.Query(CREATE_LOCKS).Exec(); err != nil {
		log.Fatalln("create transactions table fails with: ", err)
	}
}

func loadClients() {
	session := config.GetScyllaSession(config.KEYSPACE)
	defer session.Close()
	limits := []model.MonetaryValue{100000, 80000, 1000000, 10000000, 500000}
	for i, l := range limits {
		c := model.Client{
			ID:             (i + 1),
			AccountLimit:   l,
			AccountBalance: 0,
		}
		q := session.Query(INSERT_CLIENT, c.ID, c.AccountLimit, c.AccountBalance)
		if err := q.Exec(); err != nil {
			log.Fatalln("load clients fails with: ", err)
		}
	}
}
