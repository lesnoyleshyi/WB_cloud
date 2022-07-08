package HttpAdapter

import (
	"WB_cloud/internal/adapters/HttpAdapter/utils"
	"WB_cloud/internal/domain/entities"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a HttpAdapter) routeAccounts() http.Handler {
	r := chi.NewRouter()

	r.Post("/account", a.createAccount)
	r.Post("/transfer", a.transfer)
	r.Get("/balance", a.getBalance)

	return r
}

func (a HttpAdapter) createAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	account, err := utils.GetAccount(r.Body)
	if err != nil {
		a.respondError(w, "can't read body", http.StatusBadRequest, err)
		return
	}

	if err := a.accounts.Create(ctx, *account); err != nil {
		a.respondError(w, "can't create account", http.StatusInternalServerError, err)
		return
	}

	a.respondSuccess(w, "account created", http.StatusCreated)
}

func (a HttpAdapter) transfer(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	fromId, toId, amount, err := utils.GetTransferParams(r.Body)
	if err != nil {
		a.respondError(w, "can't read body", http.StatusBadRequest, err)
		return
	}

	from := entities.Account{Id: fromId, Balance: 0}
	to := entities.Account{Id: toId, Balance: 0}

	if err := a.accounts.Transfer(ctx, from, to, amount); err != nil {
		a.respondError(w, "can't make transfer", http.StatusInternalServerError, err)
		return
	}

	a.respondSuccess(w, "transfer is made successfully", http.StatusOK)
}

func (a HttpAdapter) getBalance(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	id := r.URL.Query().Get("id")
	if id == "" {
		a.respondError(w, "provide account id as query parameter: ?id=",
			http.StatusBadRequest, fmt.Errorf("bad request: empty id"))
		return
	}

	account := entities.Account{Id: id, Balance: 0}

	balance, err := a.accounts.GetBalance(ctx, account)
	if err != nil {
		a.respondError(w, fmt.Sprintf("can't get balance for account with id %s", id),
			http.StatusInternalServerError, err)
		return
	}

	a.respondSuccess(w, fmt.Sprintf("account's balance with id %s is %d", id, balance),
		http.StatusOK)
}
