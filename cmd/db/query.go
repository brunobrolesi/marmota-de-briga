package db

const CREATE_KEYSPACE = "CREATE KEYSPACE IF NOT EXISTS rinha WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'replication_factor' : 1  } AND DURABLE_WRITES = true;"

const CREATE_CLIENTS = `
CREATE TABLE IF NOT EXISTS clients (
	id int,
	account_limit int,
	account_balance int,
PRIMARY KEY (id)
);
`

const CREATE_TRANSACTIONS = `
CREATE TABLE IF NOT EXISTS transactions (
	client_id int,
	value int,
	type text,
	description text,
	created_at timestamp,
PRIMARY KEY (client_id, created_at))
WITH CLUSTERING ORDER BY (created_at DESC);
`

const CREATE_LOCKS = `
CREATE TABLE IF NOT EXISTS locks (
	client_id int,
	request_id text,
	created_at timestamp,
PRIMARY KEY (client_id, created_at, request_id))
WITH CLUSTERING ORDER BY (created_at ASC) AND default_time_to_live = 5;
`

const INSERT_CLIENT = "INSERT INTO clients (id, account_limit, account_balance) VALUES (?, ?, ?) IF NOT EXISTS"
