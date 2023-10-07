package types

import "time"

type User struct {
	balance  int
	uniqueID string
}
type Transaction struct {
	Points int       `json:"points"`
	Payer  string    `json:"payer"`
	Time   time.Time `json:"timestamp"`
}
