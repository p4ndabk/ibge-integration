package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/internal/ibge"
	"net/http"
	"strconv"
)

func SolarEfficiencieByCodeRequest(w http.ResponseWriter, r *http.Request) {
	cityId, err := strconv.Atoi(mux.Vars(r)["city_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	solarEfficiencie, erro := ibge.EfficiencieByIBGECode(cityId)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(solarEfficiencie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
