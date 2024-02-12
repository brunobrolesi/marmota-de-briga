package model

import (
	"time"
)

type TransactionID = [16]byte

type Transaction struct {
	ID          TransactionID
	ClientID    ClientID
	Value       MonetaryValue
	Type        TransactionType
	Description string
	CreatedAt   time.Time
}
