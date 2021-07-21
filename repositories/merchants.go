package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"simpl.com/repositories/models"
	. "simpl.com/loggers"
)

type MerchantsRepository interface {
	GetMerchants(whereClause string, args... interface{}) (models.MerchantSlice, error)
	PutMerchant(merchant *models.Merchant, tx *sql.Tx) error
	UpdateMerchant(merchant *models.Merchant, tx *sql.Tx) (int64, error)
	GetAllStats(merchantID string) (StatisticsPropagator, error)
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
func (repo merchantsRepositoryImpl) GetAllStats(merchantID string) (StatisticsPropagator, error) {
	var stats []totalStatistics
	var propagator StatisticsPropagator = StatisticsPropagator{}

	query := "SELECT count(*) as total_transaction_count, SUM(discount_amount) as total_discount_amount, SUM(paid_amount) as total_paid_amount, SUM(total_amount) as total_transaction_amount from transactions where merchant_id = " + merchantID
	if err := queries.Raw(query).Bind(context.Background(), repo.database, &stats); err != nil {
		Logger.Error(err)
		return propagator, err
	}

	for _, stat := range stats{
		propagator.fromTotalStatisticts(&stat)
		return propagator, nil
	}

	return propagator, errors.New("somewhing went wrong")
}


/**
Helper structs
**/
type totalStatistics struct {
	TotalTransactionCount     		sql.NullInt64  `json:"total_transaction_count" boil:"total_transaction_count"`
	TotalTransactionAmount          sql.NullFloat64  `json:"total_transaction_amount" boil:"total_transaction_amount"`
	TotalDiscountAmount          	sql.NullFloat64 `json:"total_discount_amount" boil:"total_discount_amount"` 
	TotalPaidAmount     			sql.NullFloat64  `json:"total_paid_amount" boil:"total_paid_amount"`
}

type StatisticsPropagator struct {
	TotalTransactionCount     		int64  		`json:"total_transaction_count" boil:"total_transaction_count"`
	TotalTransactionAmount          float64     `json:"total_transaction_amount" boil:"total_transaction_amount"`
	TotalDiscountAmount          	float64		`json:"total_discount_amount" boil:"total_discount_amount"` 
	TotalPaidAmount     			float64     `json:"total_paid_amount" boil:"total_paid_amount"`
}
func (statisticsPropagator *StatisticsPropagator) fromTotalStatisticts(stats *totalStatistics) {
	statisticsPropagator.TotalTransactionCount = stats.TotalTransactionCount.Int64
	statisticsPropagator.TotalDiscountAmount = stats.TotalDiscountAmount.Float64
	statisticsPropagator.TotalPaidAmount = stats.TotalPaidAmount.Float64
	statisticsPropagator.TotalTransactionAmount = stats.TotalTransactionAmount.Float64
}