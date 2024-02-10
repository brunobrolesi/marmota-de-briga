package model

type TransactionType string

const (
	Debit  TransactionType = "d"
	Credit TransactionType = "c"
)
