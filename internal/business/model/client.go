package model

type ClientID = int

type Client struct {
	ID             ClientID
	AccountLimit   MonetaryValue
	AccountBalance MonetaryValue
}

func (c *Client) GetBalanceAfterTransaction(transactionValue MonetaryValue, transactionType TransactionType) (MonetaryValue, error) {
	if transactionType == Credit {
		balanceAfterTransaction := c.AccountBalance + transactionValue
		return balanceAfterTransaction, nil
	}

	balanceAfterTransaction := c.AccountBalance - transactionValue

	// Get the absolute value of the balance result
	if balanceAfterTransaction < 0 {
		balanceAfterTransaction = -balanceAfterTransaction
	}

	if balanceAfterTransaction > c.AccountLimit {
		return c.AccountBalance, ErrClientLimitExceeded
	}

	return (c.AccountBalance - transactionValue), nil
}
