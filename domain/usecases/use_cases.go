package usecases

import (
	"context"
	"github.com/dreis0/rinha-backend/domain"
	"github.com/dreis0/rinha-backend/domain/entities"
)

type useCasesImpl struct {
	dal domain.Dal
}

func New(dal domain.Dal) domain.Usecases {
	return &useCasesImpl{
		dal: dal,
	}
}

func (u *useCasesImpl) GetStatement(ctx context.Context, customerID int) (*entities.Statement, error) {
	return u.dal.GetStatement(ctx, customerID)
}

func (u *useCasesImpl) DoTransaction(ctx context.Context, customerID int, transaction entities.Transaction) (*entities.Transaction, error) {
	return u.dal.DoTransaction(ctx, customerID, transaction)
}
