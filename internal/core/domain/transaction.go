package domain

import (
	"errors"
	"fmt"
)

var (
	ErrTransactionValidation = errors.New("validation error")
)

const (
	TransactionTypeCredit = "c"
	TransactionTypeDebit  = "d"
)

type Transaction struct {
	ID          int64
	Type        string
	Amount      int64
	Description string
	CreatedAt   int64
}

func (t Transaction) Validate() error {
	err := func(m string) error {
		return fmt.Errorf("%w: %s", ErrTransactionValidation, m)
	}

	if t.Type != TransactionTypeCredit && t.Type != TransactionTypeDebit {
		return err("type must be credit or debit")
	}

	if t.Amount < 0 {
		return err("amount must be greater than zero")
	}

	ld := len([]rune(t.Description))
	if ld < 1 || ld > 10 {
		return err("description must be between 1 and 10 characters")
	}

	return nil
}
