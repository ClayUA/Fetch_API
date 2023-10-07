package types

import "sync"

type Transaction struct {
	Points    int    `json:"points"`
	Payer     string `json:"payer"`
	Timestamp string `json:"timestamp"`
}

// slice to keep track of every transaction for this users transaction history
var TransactionHistory []Transaction

// map of all payers balances
//upated every time the balance is changed
// this is for quick retrieval when we have a GET requst

var PayerBalances = make(map[string]int)
var PayerMutex sync.Mutex
