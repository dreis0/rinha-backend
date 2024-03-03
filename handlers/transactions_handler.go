package handlers

import (
	"encoding/json"
	"errors"
	"github.com/dreis0/rinha-backend/domain"
	"github.com/dreis0/rinha-backend/domain/entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TransactionsHandler struct {
	usecase domain.Usecases
}

func NewTransactionsHandler(usecase domain.Usecases) *TransactionsHandler {
	return &TransactionsHandler{usecase: usecase}
}

type transactionRequestBody struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

func (th *TransactionsHandler) DoTransaction(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	customerIDParam := vars["id"]
	customerID, err := strconv.Atoi(customerIDParam)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	body := &transactionRequestBody{}
	err = json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	transaction := entities.Transaction{
		Value:       body.Valor,
		Type:        body.Tipo,
		Description: body.Descricao,
	}

	t, err := th.usecase.DoTransaction(req.Context(), customerID, transaction)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			res.WriteHeader(http.StatusNotFound)
		case errors.Is(err, domain.ErrNotAllowed):
			res.WriteHeader(http.StatusUnprocessableEntity)
		default:
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(res).Encode(t)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
