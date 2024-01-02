package address

import (
	"database/sql"

	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/domain/address/usecase"
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/infra/database/postgres/addressdb"
)

import (
	_ "github.com/lib/pq"
)

func NewCreateAddresUsecase(db *sql.DB) *usecase.CreateAddressUseCase {
	addressRepository := addressdb.NewAddressRepositoryDb(db)
	createAddressUseCase := usecase.NewCreateAddressUseCase(addressRepository)
	return createAddressUseCase
}
