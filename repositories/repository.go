package repositories

import (
	db "simpl.com/databases"
)

type repository struct {
	MerchantsRepository 	MerchantsRepository
	UsersRepository     	UsersRepository
	TransactionsRepository  TransactionsRepository
}

var Repositories = repository{
	MerchantsRepository: NewMerchantsRepository(db.GetConnection()),
	UsersRepository: NewUsersRepository(db.GetConnection()),
	TransactionsRepository: NewTransactionsRepository(db.GetConnection()),
}
