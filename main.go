package main

import (
	"time"
)

type User struct {
	balance  int
	uniqueID string
}
type Transaction struct {
	points int       `json:"points"`
	payer  string    `json:"payer"`
	time   time.Time `json:"timestamp"`
}

func main() {

}
