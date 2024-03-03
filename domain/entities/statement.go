package entities

import "time"

type Statement struct {
	Date         time.Time
	Customer     Customer
	Transactions []Transaction
}
