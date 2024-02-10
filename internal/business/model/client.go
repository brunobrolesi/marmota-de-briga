package model

type ClientID = int

type Client struct {
	ID      ClientID
	Limit   MonetaryValue
	Balance ClientBalance
}
