package configuration

import (
	"github.com/nosilex/crebito/internal/core/usecase"
	"github.com/nosilex/crebito/internal/infrastructure/service"
)

type Service struct {
	Account *service.AccountService
}

func NewService(ports *Port) *Service {
	accountUseCase := usecase.NewAccountUseCase(ports.accountRepository)

	return &Service{
		Account: service.NewAccountService(accountUseCase),
	}
}
