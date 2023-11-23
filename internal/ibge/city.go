package ibge

import (
	"encoding/json"
	"os"
)

type City struct {
	CodeIBGE  int     `json:"code_ibge"`
	CodeUF    int     `json:"code_uf"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Capital   int     `json:"capital"`
}

type Cities struct {
	Cities []City `json:"cities"`
}

func CityByIBGE(cityId int) (City, error) {
	jsonData, err := os.Open("storage/cities.json")
	if err != nil {
		return City{}, err
	}

	var cityData Cities
	decoder := json.NewDecoder(jsonData)
	if err := decoder.Decode(&cityData); err != nil {
		return City{}, err
	}

	var city City

	for _, c := range cityData.Cities {
		if c.CodeIBGE == cityId {
			city = c
			break
		}
	}

	return city, nil
}
