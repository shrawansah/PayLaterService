package repositories

import (
	db "simpl.com/databases"
)

type repository struct {
	MerchantsRepository MerchantsRepository
	UsersRepository     UsersRepository
}

var Repositories = repository{
	MerchantsRepository: NewMerchantsRepository(db.GetConnection()),
	UsersRepository: NewUsersRepository(db.GetConnection()),
}
