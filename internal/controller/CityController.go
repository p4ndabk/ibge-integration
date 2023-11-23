package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/internal/ibge"
	"net/http"
	"os"
	"strconv"
)

func AllCityRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := os.Open("storage/cities.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var cityData ibge.Cities
	decoder := json.NewDecoder(jsonData)
	if err := decoder.Decode(&cityData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cityData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer jsonData.Close()
}

func CityRequest(w http.ResponseWriter, r *http.Request) {
	cityId, err := strconv.Atoi(mux.Vars(r)["city_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	city, err := ibge.CityByIBGE(cityId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
