package usecase

import (
	"context"
	"github.com/nosilex/crebito/internal/core/domain"
	"github.com/nosilex/crebito/internal/core/port"
	"github.com/nosilex/crebito/pkg/web"
)

type AccountUseCase interface {
	Get(ctx context.Context, accountID int) (domain.Account, error)
	Movement(ctx context.Context, accountID int, transaction domain.Transaction) (domain.Account, error)
	Transactions(ctx context.Context, accountID int, pageable web.Pageable) (domain.AccountTransactions, error)
}

func NewAccountUseCase(accountRepository port.AccountRepository) AccountUseCase {
	return accountUseCase{
		accountRepository: accountRepository,
	}
}

type accountUseCase struct {
	accountRepository port.AccountRepository
}

func (uc accountUseCase) Get(ctx context.Context, accountID int) (domain.Account, error) {
	account, err := uc.accountRepository.Find(ctx, accountID)
	if err != nil {
		return domain.Account{}, err
	}

	return account, nil
}

func (uc accountUseCase) Movement(ctx context.Context, accountID int, transaction domain.Transaction) (domain.Account, error) {
	account, err := uc.accountRepository.Movement(ctx, accountID, transaction)
	if err != nil {
		return domain.Account{}, err
	}

	return account, nil
}

func (uc accountUseCase) Transactions(ctx context.Context, accountID int, pageable web.Pageable) (domain.AccountTransactions, error) {
	accountTransactions, err := uc.accountRepository.Transactions(ctx, accountID, pageable)
	if err != nil {
		return domain.AccountTransactions{}, err
	}

	return accountTransactions, nil
}
