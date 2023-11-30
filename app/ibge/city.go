package ibge

import (
	"github.com/p4ndabk/ibge-integration/infra/database"
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

func CityByIBGE(id int) (City, error) {
	var city City
	
	db, err := database.InitDB()
	if err != nil {
		return City{}, err
	}

	err = db.QueryRow("SELECT * FROM cities WHERE code_ibge = ?", id).Scan(&city.CodeIBGE, &city.CodeUF, &city.Name, &city.Latitude, &city.Longitude, &city.Capital)
	if err != nil {
		return City{}, err
	}

	return city, nil
}

func AllCiteis() (Cities, error) {
	var cities Cities

	rows, err := database.Query("select * from cities;")
	if err != nil {
		return Cities{}, err
	}

	for rows.Next() {
		var city City
		err = rows.Scan(&city.CodeIBGE, &city.CodeUF, &city.Name, &city.Latitude, &city.Longitude, &city.Capital)
		if err != nil {
			return Cities{}, err
		}
		cities.AddCity(city)
	}

	return cities, nil
}

func (c *Cities) AddCity(city City) {
	c.Cities = append(c.Cities, city)
}