package usecase

import (
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/domain/address/entity"
	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/domain/address/repository"
)

type CreateAddressInputDTO struct {
	ParkingPlaceID string `json:"parking_place_id"`
	Street         string `json:"street"`
	AddressNumber  int    `json:"address_number"`
	Complement     string `json:"complement"`
	District       string `json:"district"`
	UF             string `json:"uf"`
	CEP            string `json:"cep"`
	Lat            string `json:"lat"`
	Lng            string `json:"lng"`
}

type CreateAddressOutputDTO struct {
	AddressID      string `json:"address_id"`
	ParkingPlaceID string `json:"parking_place_id"`
	Street         string `json:"street"`
	AddressNumber  int    `json:"address_number"`
	Complement     string `json:"complement"`
	District       string `json:"district"`
	UF             string `json:"uf"`
	CEP            string `json:"cep"`
	Lat            string `json:"lat"`
	Lng            string `json:"lng"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}

type CreateAddressUseCase struct {
	AddressRepository repository.AddressRepository
}

func NewCreateAddressUseCase(addressRepository repository.AddressRepository) *CreateAddressUseCase {
	return &CreateAddressUseCase{
		AddressRepository: addressRepository,
	}
}

func (uc *CreateAddressUseCase) Execute(input CreateAddressInputDTO) (*CreateAddressOutputDTO, error) {
	address, err := entity.NewAddress(
		input.ParkingPlaceID,
		input.Street,
		input.AddressNumber,
		input.Complement,
		input.District,
		input.UF,
		input.CEP,
		input.Lat,
		input.Lng,
	)
	if err != nil {
		return nil, err
	}
	err = uc.AddressRepository.Save(address)
	if err != nil {
		return nil, err
	}
	return &CreateAddressOutputDTO{
		AddressID: address.AddressID,
		ParkingPlaceID: address.ParkingPlaceID,
		Street: address.Street,
		AddressNumber: address.AddressNumber,
		Complement: address.Complement,
		District: address.District,
		UF: address.UF,
		CEP: address.CEP,
		Lat: address.Lat,
		Lng: address.Lng,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
		DeletedAt: address.DeletedAt,
	}, nil
}