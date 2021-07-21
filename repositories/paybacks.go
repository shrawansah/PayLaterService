package repositories

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"simpl.com/repositories/models"
)

type PaybacksRepository interface {
	GetPaybacks(whereClause string, args... interface{}) (models.PaybackSlice, error)
	PutPayback(payback *models.Payback, tx *sql.Tx) error
	UpdatePayback(payback *models.Payback, tx *sql.Tx) (int64, error)
}

type paybacksRepositoryImpl struct {
	database *sql.DB
}

func NewPaybacksRepository(database *sql.DB) PaybacksRepository {
	return paybacksRepositoryImpl{database: database}
}

func (repo paybacksRepositoryImpl) GetPaybacks(whereClause string, args... interface{}) (models.PaybackSlice, error) {
	return models.Paybacks(qm.Where(whereClause, args...)).All(context.Background(), repo.database)
}
func (repo paybacksRepositoryImpl) PutPayback(payback *models.Payback, tx *sql.Tx) error {

	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return payback.Insert(context.Background(), contextExecutor, boil.Infer())
}
func (repo paybacksRepositoryImpl) UpdatePayback(payback *models.Payback, tx *sql.Tx) (int64, error) {
	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return payback.Update(context.Background(), contextExecutor, boil.Infer())
}