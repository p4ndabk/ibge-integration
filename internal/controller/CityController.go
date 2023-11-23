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
	var cityData ibge.Cities

	jsonData, err := os.Open("storage/cities.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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
