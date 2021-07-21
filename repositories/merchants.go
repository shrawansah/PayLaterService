package repositories

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	. "simpl.com/loggers"
	"simpl.com/repositories/models"
)

type MerchantsRepository interface {
	GetMerchants(whereClause string, args... interface{}) (models.MerchantSlice, error)
	PutMerchant(merchant *models.Merchant, tx *sql.Tx) error
	UpdateMerchant(merchant *models.Merchant, tx *sql.Tx) (int64, error)
	GetAllStats(merchantID string) (map[int64]map[string]int64, error)
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
func (repo merchantsRepositoryImpl) GetAllStats(merchantID string) (map[int64]map[string]int64, error) {
	var stats []totalStatistics
	var propagator = make(map[int64]map[string]int64)

	query := "SELECT merchant_id, count(*) as total_transaction_count, SUM(discount_amount) as total_discount_amount, SUM(paid_amount) as total_paid_amount, SUM(total_amount) as total_transaction_amount from transactions group by merchant_id "
	if merchantID != "" {
		query = "SELECT merchant_id, count(*) as total_transaction_count, SUM(discount_amount) as total_discount_amount, SUM(paid_amount) as total_paid_amount, SUM(total_amount) as total_transaction_amount from transactions where merchant_id = " + merchantID
	}
	if err := queries.Raw(query).Bind(context.Background(), repo.database, &stats); err != nil {
		Logger.Error(err)
		return propagator, err
	}

	for _, stat := range stats {
		var value = make(map[string]int64)

		value["total_transaction_count"] = stat.TotalTransactionCount.Int64
		value["total_transaction_amount"] = stat.TotalTransactionAmount.Int64
		value["total_discount_amount"] = stat.TotalDiscountAmount.Int64
		value["total_paid_amount"] = stat.TotalPaidAmount.Int64

		propagator[stat.MerchantID.Int64] = value
	}

	return propagator, nil
}


/**
Helper structs
**/
type totalStatistics struct {
	MerchantID     					sql.NullInt64  `json:"merchant_id" boil:"merchant_id"`
	TotalTransactionCount     		sql.NullInt64  `json:"total_transaction_count" boil:"total_transaction_count"`
	TotalTransactionAmount          sql.NullInt64  `json:"total_transaction_amount" boil:"total_transaction_amount"`
	TotalDiscountAmount          	sql.NullInt64 `json:"total_discount_amount" boil:"total_discount_amount"` 
	TotalPaidAmount     			sql.NullInt64  `json:"total_paid_amount" boil:"total_paid_amount"`
}