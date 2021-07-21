package repositories

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"simpl.com/repositories/models"
)

type TransactionsRepository interface {
	GetTransactions(whereClause string, args... interface{}) (models.TransactionSlice, error)
	PutTransaction(transaction *models.Transaction, tx *sql.Tx) error
	UpdateTransaction(transaction *models.Transaction, tx *sql.Tx) (int64, error)
}

type transactionsRepositoryImpl struct {
	database *sql.DB
}

func NewTransactionsRepository(database *sql.DB) TransactionsRepository {
	return transactionsRepositoryImpl{database: database}
}

func (repo transactionsRepositoryImpl) GetTransactions(whereClause string, args... interface{}) (models.TransactionSlice, error) {
	return models.Transactions(qm.Where(whereClause, args...)).All(context.Background(), repo.database)
}
func (repo transactionsRepositoryImpl) PutTransaction(transaction *models.Transaction, tx *sql.Tx) error {

	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return transaction.Insert(context.Background(), contextExecutor, boil.Infer())
}
func (repo transactionsRepositoryImpl) UpdateTransaction(transaction *models.Transaction, tx *sql.Tx) (int64, error) {
	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return transaction.Update(context.Background(), contextExecutor, boil.Infer())
}