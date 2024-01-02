package addressdb

import "database/sql"

type AddressRepositoryDb struct {
	Db *sql.DB
}

func NewAddressRepositoryDb(db *sql.DB) *AddressRepositoryDb{
	return &AddressRepositoryDb{Db: db}
}