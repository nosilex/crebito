package port

import (
	"context"

	"github.com/nosilex/crebito/internal/core/domain"
	"github.com/nosilex/crebito/pkg/web"
)

type AccountRepository interface {
	Find(ctx context.Context, id int) (domain.Account, error)
	Transactions(ctx context.Context, accountID int, pageable web.Pageable) (domain.AccountTransactions, error)
	Movement(ctx context.Context, accountID int, transaction domain.Transaction) (domain.Account, error)
}
