package domain

import (
	"errors"
)

var (
	ErrAccountNotFound          = errors.New("account: not found")
	ErrAccountInsufficientFunds = errors.New("account: insufficient funds")
)

type Account struct {
	ID         int
	HolderName string
	Limit      int64
	Balance    int64
}

type AccountTransactions struct {
	Account
	Transactions []Transaction
}
