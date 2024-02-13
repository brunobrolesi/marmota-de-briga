package model

import "time"

const TRANSACTIONS_LIMIT uint = 10

type BankStatementBalance struct {
	Total     MonetaryValue `json:"total"`
	CreatedAt time.Time     `json:"data_extrato"`
	Limit     MonetaryValue `json:"limite"`
}

type BankStatementTransaction struct {
	Value       MonetaryValue   `json:"valor"`
	Type        TransactionType `json:"tipo"`
	Description string          `json:"descricao"`
	CreatedAt   time.Time       `json:"realizada_em"`
}

type BankStatement struct {
	Balance      BankStatementBalance       `json:"saldo"`
	Transactions []BankStatementTransaction `json:"ultimas_transacoes"`
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
