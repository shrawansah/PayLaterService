package repositories

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"simpl.com/configs"
	. "simpl.com/loggers"
	"simpl.com/repositories/models"
)

type UsersRepository interface {
	GetUsers(whereClause string, args... interface{}) (models.UserSlice, error)
	PutUser(user *models.User, tx *sql.Tx) error
	UpdateUser(user *models.User, tx *sql.Tx) (int64, error)
	GetAllStats(userID string) (UserStatistics, error)
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
func (repo usersRepositoryImpl) GetAllStats(userID string) (UserStatistics, error) {

	var stats []userTotalStatistics
	var users = make(map[int64]map[string]int64)
	var usersAtLimit = []int64{}
	var usersAtLimitMap = make(map[int64]bool)

	var propagator = UserStatistics{
		Users: 			users,
		UsersAtLimit: 	usersAtLimit,
		TotalDueAmount: 0,
	}

	query := "select transactions.id as transaction_id, due_amount, credit_limit, user_id, total_amount, discount_amount, paid_amount from users inner join transactions on users.id = transactions.user_id "

	if err := queries.Raw(query).Bind(context.Background(), repo.database, &stats); err != nil {
		Logger.Error(err)
		return propagator, err
	}

	for _, stat := range stats {
		var userID = stat.UserID.Int64
		if _, ok := propagator.Users[userID]; !ok {
			propagator.Users[userID] = make(map[string]int64)
		} 
		if _, ok := propagator.Users[userID]["total_transaction_amount"]; !ok {
			propagator.Users[userID]["total_transaction_amount"] = 0
		}
		if _, ok := propagator.Users[userID]["total_transactions_count"]; !ok {
			propagator.Users[userID]["total_transactions_count"] = 0
		}
		if _, ok := propagator.Users[userID]["total_paid_amount"]; !ok {
			propagator.Users[userID]["total_paid_amount"] = 0
		}
		if _, ok := propagator.Users[userID]["total_discount_amount"]; !ok {
			propagator.Users[userID]["total_discount_amount"] = 0
		}

		propagator.Users[userID]["total_transactions_count"]++
		propagator.Users[userID]["total_due_amount"] = stat.TotalDueAmount.Int64
		propagator.Users[userID]["total_transaction_amount"] += stat.TotalTransactionAmount.Int64
		propagator.Users[userID]["total_paid_amount"] += stat.TotalPaidAmount.Int64
		propagator.Users[userID]["total_discount_amount"] += stat.TotalDiscountAmount.Int64
		propagator.Users[userID]["total_credit_limit"] = stat.UserCreditLimit.Int64

		if stat.UserCreditLimit.Int64 - stat.TotalDueAmount.Int64 <= configs.Configs.UserAtCreditLimitThreshold {
			usersAtLimitMap[userID] = true
		}
	}

	for key, _ := range usersAtLimitMap {
		propagator.UsersAtLimit = append(propagator.UsersAtLimit, key)
	}

	userIDInt64, _ := strconv.ParseInt(userID, 10, 64)
	for key, userData := range propagator.Users {
		propagator.TotalDueAmount += userData["total_due_amount"]
		if userID != "" && key != userIDInt64 {
			delete(propagator.Users , key);
		}
	}

	return propagator, nil
}


/**
Helper structs
**/
type userTotalStatistics struct {
	TransactionID     				sql.NullInt64  `json:"transaction_id" boil:"transaction_id"`
	TotalTransactionAmount          sql.NullInt64  `json:"total_amount" boil:"total_amount"`
	TotalDiscountAmount          	sql.NullInt64  `json:"discount_amount" boil:"discount_amount"` 
	TotalPaidAmount     			sql.NullInt64  `json:"paid_amount" boil:"paid_amount"`
	TotalDueAmount					sql.NullInt64  `json:"due_amount" boil:"due_amount"`
	UserID							sql.NullInt64  `json:"user_id" boil:"user_id"`
	UserCreditLimit					sql.NullInt64  `json:"credit_limit" boil:"credit_limit"`
}

type UserStatistics struct {
	Users    						map[int64]map[string]int64    `json:"users" boil:"users"`
	UsersAtLimit					[]int64						  `json:"users_at_credit_limit" boil:"users_at_credit_limit"`
	TotalDueAmount					int64						  `json:"total_due_amount" boil:"total_due_amount"`
}