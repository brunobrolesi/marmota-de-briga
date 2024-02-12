package model

type ClientBalance MonetaryValue

func (c *ClientBalance) CanReceiveTransaction(transactionValue MonetaryValue, clientLimit MonetaryValue, transactionType TransactionType) bool {
	if transactionType == Credit {
		return true
	}

	balanceAfterTransaction := MonetaryValue(*c) - transactionValue

	// Get the absolute value of the balance result
	if balanceAfterTransaction < 0 {
		balanceAfterTransaction = -balanceAfterTransaction
	}

	return balanceAfterTransaction < MonetaryValue(clientLimit)
}

func (c *ClientBalance) CanNotReceiveTransaction(transactionValue MonetaryValue, clientLimit MonetaryValue, transactionType TransactionType) bool {
	return !c.CanReceiveTransaction(transactionValue, clientLimit, transactionType)
}

func (c *ClientBalance) AddTransaction(transactionValue MonetaryValue, transactionType TransactionType) {
	if transactionType == Credit {
		*c = ClientBalance(MonetaryValue(*c) + transactionValue)
		return
	}

	*c = ClientBalance(MonetaryValue(*c) - transactionValue)

}
