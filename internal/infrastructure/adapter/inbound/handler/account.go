package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/nosilex/crebito/internal/core/domain"
	"github.com/nosilex/crebito/internal/infrastructure/adapter/inbound/handler/dto"
	"github.com/nosilex/crebito/internal/infrastructure/service"
	"github.com/nosilex/crebito/pkg/web"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h AccountHandler) Movement(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accountID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		web.Response.Error(w, errors.New("invalid account id"), http.StatusBadRequest)
		return
	}

	transactionRequest := dto.TransactionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&transactionRequest); err != nil {
		web.Response.Error(w, err, http.StatusBadRequest)
		return
	}

	account, err := h.accountService.Movement(ctx, accountID, transactionRequest.MapToDomain())
	if err != nil {
		if errors.Is(err, domain.ErrTransactionValidation) || errors.Is(err, domain.ErrAccountInsufficientFunds) {
			web.Response.Error(w, err, http.StatusUnprocessableEntity)
			return
		}

		if errors.Is(err, domain.ErrAccountNotFound) {
			web.Response.Error(w, err, http.StatusNotFound)
			return
		}

		web.Response.Error(w, err, http.StatusInternalServerError)
		return
	}

	web.Response.OK(w, dto.TransactionResponse{
		Limit:   account.Limit,
		Balance: account.Balance,
	})
}

func (h AccountHandler) Transactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accountID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		web.Response.Error(w, errors.New("invalid account id"), http.StatusBadRequest)
		return
	}

	accountTransactions, err := h.accountService.Transactions(ctx, accountID, web.NewPageable(1, 10))
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			web.Response.Error(w, err, http.StatusNotFound)
			return
		}

		web.Response.Error(w, err, http.StatusInternalServerError)
		return
	}

	web.Response.OK(w, new(dto.AccountResponse).MapFromDomain(accountTransactions))
}
