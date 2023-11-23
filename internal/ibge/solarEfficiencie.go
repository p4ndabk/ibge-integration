package ibge

import (
	"encoding/json"
	"os"
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

func EfficiencieByIBGECode(cityId int) (SolarEfficiencie, error) {
	var err error
	s := SolarEfficiencies{}
	s, err = s.GetSolarEfficiencies()
	if err != nil {
		return SolarEfficiencie{}, err
	}

	city, err := CityByIBGE(cityId)
	if err != nil {
		return SolarEfficiencie{}, err
	}

	var solarEfficiencie SolarEfficiencie

	for _, s := range s.SolarEfficiencies {
		if int(s.Latitude) >= int(city.Latitude) && s.Longitude >= city.Longitude {
			solarEfficiencie = s
			break
		}
	}

	if solarEfficiencie.ID == 0 {
		return SolarEfficiencie{}, err
	}

	return solarEfficiencie, nil
}
