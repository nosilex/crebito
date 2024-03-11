package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nosilex/crebito/internal/core/domain"
	"github.com/nosilex/crebito/internal/core/port"
	"github.com/nosilex/crebito/pkg/helper"
	"github.com/nosilex/crebito/pkg/web"
)

func NewAccountRepository(db *sql.DB) port.AccountRepository {
	return accountRepository{
		db: db,
	}
}

type accountRepository struct {
	db *sql.DB
}

func (a accountRepository) Find(ctx context.Context, id int) (domain.Account, error) {
	var account domain.Account
	if err := a.db.QueryRowContext(
		ctx,
		"SELECT id, holder_name, limit_amount, balance_amount FROM account WHERE id = ?", id,
	).Scan(&account.ID, &account.HolderName, &account.Limit, &account.Balance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Account{}, domain.ErrAccountNotFound
		}
		return domain.Account{}, err
	}

	return account, nil
}

func (a accountRepository) Transactions(ctx context.Context, accountID int, pageable web.Pageable) (domain.AccountTransactions, error) {
	rows, err := a.db.QueryContext(
		ctx,
		`
			SELECT
				a.limit_amount,
				a.balance_amount,
				COALESCE(t.id, 0) as id,
				COALESCE(t.type_id, '') as type_id,
				COALESCE(t.amount, 0) as amount,
				COALESCE(t.description, '') as description,
				COALESCE(t.created_at, NOW()) as created_at
			FROM
				account a
				LEFT JOIN transaction t ON a.id = t.account_id
			WHERE
				a.id = ?
			ORDER BY
				t.created_at DESC
				LIMIT ? OFFSET ?
		`,
		accountID,
		pageable.Limit(),
		pageable.Offset(),
	)
	if err != nil {
		return domain.AccountTransactions{}, err
	}
	defer rows.Close()

	rowsCount := 0
	var limit, balance int64
	transactions := make([]domain.Transaction, 0, pageable.Limit())
	for rows.Next() {
		var ID int64
		var typeID string
		var amount int64
		var description string
		var createdAt time.Time

		if err := rows.Scan(
			&limit,
			&balance,
			&ID,
			&typeID,
			&amount,
			&description,
			&createdAt,
		); err != nil {
			return domain.AccountTransactions{}, err
		}

		if ID > 0 {
			transactions = append(transactions, domain.Transaction{
				ID:          ID,
				Type:        typeID,
				Amount:      amount,
				Description: description,
				CreatedAt:   createdAt.UnixMicro(),
			})
		}
		rowsCount++
	}

	if rowsCount == 0 {
		return domain.AccountTransactions{}, domain.ErrAccountNotFound
	}

	return domain.AccountTransactions{
		Account: domain.Account{
			Limit:   limit,
			Balance: balance,
		},
		Transactions: transactions,
	}, nil
}

func (a accountRepository) Movement(ctx context.Context, accountID int, transaction domain.Transaction) (domain.Account, error) {
	amount := helper.If(transaction.Type == domain.TransactionTypeCredit, transaction.Amount, transaction.Amount*-1)

	// Let's process the transaction in DB
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Account{}, err
	}
	defer tx.Rollback()

	var account domain.Account
	var availability bool
	if err := tx.QueryRowContext(
		ctx,
		`
			SELECT HIGH_PRIORITY
				id,
				holder_name,
				limit_amount,
				balance_amount,
				( limit_amount + ( balance_amount + ?) >= 0 ) AS availability 
			FROM
				account 
			WHERE
				id = ? FOR UPDATE
		`,
		amount, accountID,
	).Scan(&account.ID, &account.HolderName, &account.Limit, &account.Balance, &availability); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Account{}, domain.ErrAccountNotFound
		}

		return domain.Account{}, err
	}
	if !availability {
		return domain.Account{}, domain.ErrAccountInsufficientFunds
	}

	account.Balance += amount
	if _, err := tx.ExecContext(
		ctx,
		"UPDATE account SET balance_amount = ? WHERE id = ?",
		account.Balance,
		account.ID,
	); err != nil {
		return domain.Account{}, err
	}

	if _, err := tx.ExecContext(
		ctx,
		`
			INSERT LOW_PRIORITY INTO transaction ( account_id, type_id, amount, description )
			VALUES (?, ?, ?, ?)
		`,
		accountID,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
	); err != nil {
		return domain.Account{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.Account{}, err
	}

	return account, nil
}
