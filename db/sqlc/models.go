// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Account struct {
	ID        int32
	Username  string
	Ammount   int64
	Currency  string
	CreatedAt time.Time
}

type Entry struct {
	ID int32
	// can be negative
	Amount    string
	AccountID int32
	CreatedAt time.Time
}

type Transaction struct {
	ID            int32
	FromAccountID int32
	ToAccountID   int32
	// Have be positive
	Amount    int64
	CreatedAt time.Time
}
