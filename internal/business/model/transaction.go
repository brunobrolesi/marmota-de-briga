package model

import "time"

type TransactionID = int

type Transaction struct {
	ID          TransactionID
	ClientID    ClientID
	Value       MonetaryValue
	Type        TransactionType
	Description string
	CreatedAt   time.Time
}
