package config

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

const KEYSPACE = "rinha"

func NewScyllaClient(keyspace string) *gocqlx.Session {
	var err error
	var session gocqlx.Session
	for ok := true; ok; ok = err != nil {
		cluster := gocql.NewCluster("scylla-db")
		cluster.Keyspace = keyspace
		session, err = gocqlx.WrapSession(cluster.CreateSession())
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
		}
	}
	return &session

}
