package address

import (
	"database/sql"

	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/infra/database/postgres/addressdb"
)

import (
	_ "github.com/lib/pq"
)

func NewWebAddressHandler(db *sql.DB) *AddressHandler {
	addressRepository := addressdb.NewAddressRepositoryDb(db)
	webAddressHandler := NewAddressHandler(addressRepository)
	return webAddressHandler
}
