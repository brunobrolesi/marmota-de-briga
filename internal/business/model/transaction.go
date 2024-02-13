package model

import (
	"time"
)

type Transaction struct {
	ClientID    ClientID
	Value       MonetaryValue
	Type        TransactionType
	Description string
	CreatedAt   time.Time
}
