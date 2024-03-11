package configuration

import (
	"fmt"

	"github.com/nosilex/crebito/internal/core/port"
	"github.com/nosilex/crebito/internal/infrastructure/adapter/outbound/repository"
)

type Port struct {
	accountRepository port.AccountRepository
}

func NewPort() (*Port, error) {

	db, err := newMariaDB()
	if err != nil {
		return nil, fmt.Errorf("init database connection: %w", err)
	}

	accountRepository := repository.NewAccountRepository(db)

	return &Port{
		accountRepository: accountRepository,
	}, nil
}
