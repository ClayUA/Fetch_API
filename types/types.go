package types

import (
	"time"
)

// struct for any given transaction
// times will be read in and compared in strings then converted to time.Time values for sorting purposes
type Transaction struct {
	Points     int    `json:"points"`
	Payer      string `json:"payer"`
	Timestamp  string `json:"timestamp"`
	ParsedTime time.Time
}

// slice to keep track of every transaction for this users transaction history
var TransactionHistory []Transaction

// map of all payers balances
//upated every time the balance is changed
// this is for quick retrieval when we have a GET requst

// Gloabal map to keep track of Payer Balances so we can always call /balance and get accurate data
var PayerBalances = make(map[string]int)

// Just a quick way to validate if we can spend the requested points
var TotalPointsBalance int

// I only use this to read in the amount of points on a /spend endpoint
type SpendRequest struct {
	Points int `json:"points"`
}
