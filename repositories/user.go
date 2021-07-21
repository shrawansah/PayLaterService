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

type UsersRepository interface {
	GetUsers(whereClause string, args... interface{}) (models.UserSlice, error)
	PutUser(user *models.User, tx *sql.Tx) error
	UpdateUser(user *models.User, tx *sql.Tx) (int64, error)
	GetAllStats(userID string) (StatisticsPropagator, error)
}

type usersRepositoryImpl struct {
	database *sql.DB
}

func NewUsersRepository(database *sql.DB) UsersRepository {
	return usersRepositoryImpl{database: database}
}

func (repo usersRepositoryImpl) GetUsers(whereClause string, args... interface{}) (models.UserSlice, error) {
	return models.Users(qm.Where(whereClause, args...)).All(context.Background(), repo.database)
}
func (repo usersRepositoryImpl) PutUser(user *models.User, tx *sql.Tx) error {

	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return user.Insert(context.Background(), contextExecutor, boil.Infer())
}
func (repo usersRepositoryImpl) UpdateUser(user *models.User, tx *sql.Tx) (int64, error) {
	var contextExecutor boil.ContextExecutor
	contextExecutor = tx
	if tx == nil {
		contextExecutor = repo.database
	}

	return user.Update(context.Background(), contextExecutor, boil.Infer())
}
func (repo usersRepositoryImpl) GetAllStats(userID string) (StatisticsPropagator, error) {
	var stats []totalStatistics
	var propagator StatisticsPropagator = StatisticsPropagator{}

	query := "SELECT count(*) as total_transaction_count, SUM(discount_amount) as total_discount_amount, SUM(paid_amount) as total_paid_amount, SUM(total_amount) as total_transaction_amount from transactions where merchant_id = " + userID
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
type userTotalStatistics struct {
	TotalTransactionCount     		sql.NullInt64  `json:"total_transaction_count" boil:"total_transaction_count"`
	TotalTransactionAmount          sql.NullFloat64  `json:"total_transaction_amount" boil:"total_transaction_amount"`
	TotalDiscountAmount          	sql.NullFloat64 `json:"total_discount_amount" boil:"total_discount_amount"` 
	TotalPaidAmount     			sql.NullFloat64  `json:"total_paid_amount" boil:"total_paid_amount"`
}

type UserStatisticsPropagator struct {
	TotalTransactionCount     		int64  		`json:"total_transaction_count" boil:"total_transaction_count"`
	TotalTransactionAmount          float64     `json:"total_transaction_amount" boil:"total_transaction_amount"`
	TotalDiscountAmount          	float64		`json:"total_discount_amount" boil:"total_discount_amount"` 
	TotalPaidAmount     			float64     `json:"total_paid_amount" boil:"total_paid_amount"`
}
func (statisticsPropagator *UserStatisticsPropagator) fromTotalStatisticts(stats *userTotalStatistics) {
	statisticsPropagator.TotalTransactionCount = stats.TotalTransactionCount.Int64
	statisticsPropagator.TotalDiscountAmount = stats.TotalDiscountAmount.Float64
	statisticsPropagator.TotalPaidAmount = stats.TotalPaidAmount.Float64
	statisticsPropagator.TotalTransactionAmount = stats.TotalTransactionAmount.Float64
}