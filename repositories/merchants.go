package repositories

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"simpl.com/repositories/models"
)

type MerchantsRepository interface {
	GetMerchants(whereClause string, args... interface{}) (models.MerchantSlice, error)
	PutMerchant(merchant *models.Merchant, tx *sql.Tx) error
	UpdateMerchant(merchant *models.Merchant, tx *sql.Tx) (int64, error)
}

type merchantsRepositoryImpl struct {
	database *sql.DB
}
func NewMerchantsRepository(database *sql.DB) MerchantsRepository {
	return merchantsRepositoryImpl{database: database}
}

func (repo merchantsRepositoryImpl) GetMerchants(whereClause string, args... interface{}) (models.MerchantSlice, error) {
	return models.Merchants(qm.Where(whereClause, args...)).All(context.Background(), repo.database)
}
func (repo merchantsRepositoryImpl) PutMerchant(merchant *models.Merchant, tx *sql.Tx) error {

	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return merchant.Insert(context.Background(), contextExecutor, boil.Infer())
}
func (repo merchantsRepositoryImpl) UpdateMerchant(merchant *models.Merchant, tx *sql.Tx) (int64, error) {
	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return merchant.Update(context.Background(), contextExecutor, boil.Infer())
}