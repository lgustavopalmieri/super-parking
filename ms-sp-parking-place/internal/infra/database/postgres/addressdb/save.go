package addressdb

import "github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/domain/address/entity"

const query = `INSERT INTO address (
			address_id,
			parking_place_id,
			street,
			number,
			complement,
			district,
			uf,
			cep,
			lat,
			lng
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

func (r *AddressRepositoryDb) Save(address *entity.Address) error {
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		address.AddressID,
		address.ParkingPlaceID,
		address.Street,
		address.AddressNumber,
		address.Complement,
		address.District,
		address.UF,
		address.CEP,
		address.Lat,
		address.Lng)
	if err != nil {
		return err
	}
	return nil
}
