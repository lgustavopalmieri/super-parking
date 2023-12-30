package repository

import "github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/address/entity"

type AddressRepository interface {
	Save(address *entity.Address) error
}