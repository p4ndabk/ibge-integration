package ibge

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type SolarEfficiencie struct {
	ID        int     `json:"id"`
	Country   string  `json:"country"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Annual    int     `json:"annual"`
	January   int     `json:"january"`
	February  int     `json:"february"`
	March     int     `json:"march"`
	April     int     `json:"april"`
	May       int     `json:"may"`
	June      int     `json:"june"`
	July      int     `json:"july"`
	August    int     `json:"august"`
	September int     `json:"september"`
	October   int     `json:"october"`
	November  int     `json:"november"`
	December  int     `json:"december"`
}

type SolarEfficiencies struct {
	SolarEfficiencies []SolarEfficiencie `json:"solar_efficiencies"`
}

func (s SolarEfficiencies) GetSolarEfficiencies() (SolarEfficiencies, error) {
	jsonData, err := os.Open("storage/solar_efficiencies.json")
	if err != nil {
		return SolarEfficiencies{}, err
	}
	var solarEfficienciesData SolarEfficiencies

	decoder := json.NewDecoder(jsonData)
	if err := decoder.Decode(&solarEfficienciesData); err != nil {
		return SolarEfficiencies{}, err
	}

	return solarEfficienciesData, nil
}

func SolarEfficiencieByCodeRequest(w http.ResponseWriter, r *http.Request) {
	cityId, err := strconv.Atoi(mux.Vars(r)["city_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	solarEfficiencie, erro := GetEfficiencieByIBGECode(cityId)
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

func GetEfficiencieByIBGECode(cityId int) (SolarEfficiencie, error) {
	var erro error
	s := SolarEfficiencies{}
	s, erro = s.GetSolarEfficiencies()
	if erro != nil {
		return SolarEfficiencie{}, erro
	}

	city, err := CityByIBGE(cityId)
	if err != nil {
		return SolarEfficiencie{}, err
	}

	fmt.Println(city.Latitude, city.Longitude)

	var solarEfficiencie SolarEfficiencie

	for _, s := range s.SolarEfficiencies {
		if int(s.Latitude) >= int(city.Latitude) && s.Longitude >= city.Longitude { // substitua ID pelo campo que contém o código IBGE
			solarEfficiencie = s
			break
		}
	}

	if solarEfficiencie.ID == 0 { // substitua ID pelo campo que contém o código IBGE
		return SolarEfficiencie{}, erro
	}

	return solarEfficiencie, nil
}
