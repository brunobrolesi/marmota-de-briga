package model

type ClientID = int

type Client struct {
	ID             ClientID
	AccountLimit   MonetaryValue
	AccountBalance ClientBalance
}
