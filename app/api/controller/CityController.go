package controller

import (
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/app/helper"
	"github.com/p4ndabk/ibge-integration/app/ibge"
)

func AllCityRequest(w http.ResponseWriter, r *http.Request) {
	cityData := &ibge.Cities{}

	cityData, err := ibge.AllCiteis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cityData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CityRequest(w http.ResponseWriter, r *http.Request) {
	cityId, err := helper.StringToInt(mux.Vars(r)["city_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	city, err := ibge.CityByIBGE(cityId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CheckCoordinateRequest(w http.ResponseWriter, r *http.Request) {
	ibegeCode := mux.Vars(r)["ibge_code"]

	body, err := ibge.GetCoordinatesIBGE(ibegeCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
