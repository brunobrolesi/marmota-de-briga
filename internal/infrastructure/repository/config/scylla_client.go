package config

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

const KEYSPACE = "rinha"

func NewScyllaClient(keyspace string) *gocqlx.Session {
	cluster := gocql.NewCluster("scylla-db")
	cluster.Keyspace = keyspace
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic(err)
	}
	return &session
}
