package postgres

import (
	"context"
	"errors"
	"github.com/dreis0/rinha-backend/domain"
	"github.com/dreis0/rinha-backend/domain/entities"
	"github.com/jackc/pgx/v5"
)

type popstgresDal struct {
	conn pgx.Conn
}

var _ domain.Dal = &popstgresDal{}

func NewPostgresDal(conn pgx.Conn) domain.Dal {
	return &popstgresDal{
		conn: conn,
	}
}

func (dal *popstgresDal) GetStatement(ctx context.Context, id int) (*entities.Statement, error) {
	panic("implement me")
}

func (dal *popstgresDal) DoTransaction(ctx context.Context, customerID int, transaction entities.Transaction) (*entities.Transaction, error) {
	customer := &entities.Customer{}
	row := dal.conn.QueryRow(ctx, "select id, account_limit, initial_balance from customer where id = $1", customerID)
	err := row.Scan(&customer.ID, &customer.Limit, &customer.InitialBalance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	value := transaction.Value
	if transaction.Type == entities.Debit {
		value = -value
	}

	tx, err := dal.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	var balance int
	tx.QueryRow(ctx, "select balance from customer_balance where customer_id = $1", customerID).Scan(&balance)

	if balance+value > customer.Limit || balance+value < -customer.Limit {
		return nil, domain.ErrNotAllowed
	}

	_, err = tx.Exec(ctx,
		"insert into transaction (customer_id, value, type, description) values ($1, $2, $3, $4)",
		customerID,
		transaction.Value,
		transaction.Type,
		transaction.Description,
	)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	_, err = tx.Exec(ctx, "insert into customer_balance (customer_id, balance) "+
		"values ($1, $2) "+
		"ON CONFLICT (customer_id) "+
		"DO update set balance = customer_balance.balance + $3", customerID, value, value)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
