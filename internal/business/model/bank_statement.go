package model

import "time"

const TRANSACTIONS_LIMIT uint = 10

type BankStatementBalance struct {
	Total     MonetaryValue
	CreatedAt time.Time
	Limit     MonetaryValue
}

type BankStatementTransaction struct {
	Value       MonetaryValue
	Type        TransactionType
	Description string
	CreatedAt   time.Time
}

type BankStatement struct {
	Balance      BankStatementBalance
	Transactions []BankStatementTransaction
}

func ToBankStatementTransactions(transactions []Transaction) []BankStatementTransaction {
	var bankStatementTransactions []BankStatementTransaction
	for _, t := range transactions {
		bankStatementTransactions = append(bankStatementTransactions, BankStatementTransaction{
			Value:       t.Value,
			Type:        t.Type,
			Description: t.Description,
			CreatedAt:   t.CreatedAt,
		})
	}
	return bankStatementTransactions
}
