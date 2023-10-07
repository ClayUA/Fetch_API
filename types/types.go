package types

type User struct {
	balance  int
	uniqueID string
}
type Transaction struct {
	Points    int    `json:"points"`
	Payer     string `json:"payer"`
	Timestamp string `json:"timestamp"`
}
