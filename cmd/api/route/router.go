package route

import (
	"net/http"

	"github.com/nosilex/crebito/internal/infrastructure/configuration"
)

func Setup(handler *configuration.Handler) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("POST /clientes/{id}/transacoes", handler.Account.Movement)
	r.HandleFunc("GET /clientes/{id}/extrato", handler.Account.Transactions)

	return r
}
