package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Address struct {
	AddressID      string
	ParkingPlaceID string
	Street         string
	AddressNumber  int
	Complement     string
	District       string
	UF             string
	CEP            string
	Lat            string
	Lng            string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
}

func NewAddress(parking_place_id string, street string,
	address_number int, complement, district, uf, cep,
	lat, lng string,
) (*Address, error) {
	currentTime := time.Now()

	address := &Address{
		AddressID:      uuid.New().String(),
		ParkingPlaceID: parking_place_id,
		Street:         street,
		AddressNumber:  address_number,
		Complement:     complement,
		District:       district,
		UF:             uf,
		CEP:            cep,
		Lat:            lat,
		Lng:            lng,
		CreatedAt:      currentTime.Format("2006-01-02T15:04:05"),
		UpdatedAt:      currentTime.Format("2006-01-02T15:04:05"),
		DeletedAt:      "",
	}
	err := address.Validate()
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (a *Address) Validate() error {
	if a.AddressID == "" {
		return errors.New("invalid id")
	}
	if a.ParkingPlaceID == "" {
		return errors.New("invalid parking_place_id")
	}
	if a.Street == "" {
		return errors.New("invalid street")
	}
	if a.AddressNumber <= 0 {
		return errors.New("invalid address number")
	}
	if a.District == "" {
		return errors.New("invalid district")
	}
	if a.UF == "" {
		return errors.New("invalid uf")
	}
	if a.CEP == "" {
		return errors.New("invalid cep")
	}
	return nil
}
