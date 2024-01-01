package address

import "github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/domain/address/repository"

type AddressHandler struct {
	AddressRepository repository.AddressRepository
}

func NewAddressHandler(
	AddressRepository repository.AddressRepository,
) *AddressHandler {
	return &AddressHandler{
		AddressRepository: AddressRepository,
	}
}
