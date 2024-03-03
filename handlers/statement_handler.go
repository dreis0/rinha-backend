package handlers

import (
	"github.com/dreis0/rinha-backend/domain"
	"net/http"
)

type StatementHandler struct {
	usecase domain.Usecases
}

func NewStatementHandler(usecase domain.Usecases) *StatementHandler {
	return &StatementHandler{usecase: usecase}
}

func (sh *StatementHandler) GetCustomerStatement(res http.ResponseWriter, req *http.Request) {

}
