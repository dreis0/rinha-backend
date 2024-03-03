package domain

import (
	"context"
	"github.com/dreis0/rinha-backend/domain/entities"
)

type Usecases interface {
	GetStatement(ctx context.Context, customerID int) (*entities.Statement, error)

	DoTransaction(ctx context.Context, customerID int, transaction entities.Transaction) (*entities.Transaction, error)
}
