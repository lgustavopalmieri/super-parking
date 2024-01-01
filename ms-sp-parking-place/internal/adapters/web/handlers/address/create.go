package address

import (
	"encoding/json"
	"net/http"

	"github.com/lgustavopalmieri/super-parking/ms-sp-parking-place/internal/domain/address/usecase"
)

func (h *AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateAddressInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createAddress := usecase.NewCreateAddressUseCase(h.AddressRepository)
	output, err := createAddress.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}