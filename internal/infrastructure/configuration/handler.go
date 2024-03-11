package configuration

import "github.com/nosilex/crebito/internal/infrastructure/adapter/inbound/handler"

type Handler struct {
	Account *handler.AccountHandler
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Account: handler.NewAccountHandler(service.Account),
	}
}
