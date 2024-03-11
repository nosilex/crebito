package service

import (
	"context"

	"github.com/nosilex/crebito/internal/core/domain"
	"github.com/nosilex/crebito/internal/core/usecase"
	"github.com/nosilex/crebito/pkg/web"
)

type AccountService struct {
	accountUseCase usecase.AccountUseCase
}

func NewAccountService(usecase usecase.AccountUseCase) *AccountService {
	return &AccountService{
		accountUseCase: usecase,
	}
}

func (s AccountService) Movement(ctx context.Context, accountID int, transaction domain.Transaction) (domain.Account, error) {
	if err := transaction.Validate(); err != nil {
		return domain.Account{}, err
	}

	return s.accountUseCase.Movement(ctx, accountID, transaction)
}

func (s AccountService) Transactions(ctx context.Context, accountID int, pageable web.Pageable) (domain.AccountTransactions, error) {
	accountTransactions, err := s.accountUseCase.Transactions(ctx, accountID, pageable)
	if err != nil {
		return domain.AccountTransactions{}, err
	}

	return accountTransactions, nil
}
