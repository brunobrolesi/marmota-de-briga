package db

const (
	CREATE_KEYSPACE = "CREATE KEYSPACE IF NOT EXISTS rinha WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'replication_factor' : 1  } AND DURABLE_WRITES = true;"
)
