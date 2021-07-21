package repositories

import (
	db "simpl.com/databases"
)

type repository struct {
	MerchantsRepository 	MerchantsRepository
	UsersRepository     	UsersRepository
	TransactionsRepository  TransactionsRepository
	PaybacksRepository   	PaybacksRepository
}

var Repositories = repository{
	MerchantsRepository: NewMerchantsRepository(db.GetConnection()),
	UsersRepository: NewUsersRepository(db.GetConnection()),
	TransactionsRepository: NewTransactionsRepository(db.GetConnection()),
	PaybacksRepository: NewPaybacksRepository(db.GetConnection()),
}
