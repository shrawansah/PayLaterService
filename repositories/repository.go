package repositories

import (
	db "simpl.com/databases"
)

type repository struct {
	MerchantsRepository MerchantsRepository
}

var Repositories = repository{
	MerchantsRepository: NewMerchantsRepository(db.GetConnection()),
}
