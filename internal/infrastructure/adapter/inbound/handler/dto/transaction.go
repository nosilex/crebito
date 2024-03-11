package dto

import (
	"github.com/nosilex/crebito/internal/core/domain"
)

type TransactionRequest struct {
	Amount      int64  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func (r TransactionRequest) MapToDomain() domain.Transaction {
	return domain.Transaction{
		Type:        r.Type,
		Amount:      r.Amount,
		Description: r.Description,
	}
}

type TransactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}
