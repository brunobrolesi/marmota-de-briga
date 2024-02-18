package config

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2/log"
)

const KEYSPACE = "rinha"

func GetScyllaSession(keyspace string) *gocql.Session {
	var err error
	var session *gocql.Session
	for ok := true; ok; ok = err != nil {
		cluster := gocql.NewCluster("scylla-db")
		cluster.Keyspace = keyspace
		if err != nil {
			log.Error(err)
			time.Sleep(5 * time.Second)
		}
		cluster.NumConns = 10
		cluster.Consistency = gocql.LocalQuorum
		cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}
		session, err = cluster.CreateSession()
		if err != nil {
			log.Error("failed to create session: %v", err)
		}
	}
	return session
}
