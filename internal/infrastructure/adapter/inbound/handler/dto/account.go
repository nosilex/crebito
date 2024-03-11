package dto

import (
	"time"

	"github.com/nosilex/crebito/internal/core/domain"
)

type AccountResponse struct {
	Balance      AccountBalanceResponse        `json:"saldo"`
	Transactions []AccountTransactionsResponse `json:"ultimas_transacoes"`
}

type AccountBalanceResponse struct {
	Total    int64     `json:"total"`
	IssuedAt time.Time `json:"data_extrato"`
	Limit    int64     `json:"limite"`
}

type AccountTransactionsResponse struct {
	Amount      int64     `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

func (r AccountResponse) MapFromDomain(accountTransactions domain.AccountTransactions) AccountResponse {
	accountResponse := AccountResponse{
		Balance: AccountBalanceResponse{
			Total:    accountTransactions.Balance,
			IssuedAt: time.Now().UTC(),
			Limit:    accountTransactions.Limit,
		},
	}

	transactions := make([]AccountTransactionsResponse, 0, len(accountTransactions.Transactions))
	for _, transaction := range accountTransactions.Transactions {
		transactions = append(transactions, AccountTransactionsResponse{
			Amount:      transaction.Amount,
			Type:        transaction.Type,
			Description: transaction.Description,
			CreatedAt:   time.UnixMicro(transaction.CreatedAt).UTC(),
		})
	}
	accountResponse.Transactions = transactions

	return accountResponse
}
