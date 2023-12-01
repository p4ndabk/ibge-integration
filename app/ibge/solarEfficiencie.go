package ibge

import (
	"github.com/p4ndabk/ibge-integration/infra/database"
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

func EfficiencieByIBGECode(cityId int) (SolarEfficiencie, error) {
	var err error
	var solarEfficiencie SolarEfficiencie

	city, err := CityByIBGE(cityId)
	if err != nil {
		return SolarEfficiencie{}, err
	}

	db, err := database.InitDB()
	if err != nil {
		return solarEfficiencie, err
	}
	err = db.QueryRow("SELECT * FROM solar_efficiencies WHERE latitude >= ? AND longitude >= ?",
		city.Latitude, city.Longitude).Scan(
		&solarEfficiencie.ID, &solarEfficiencie.Country,
		&solarEfficiencie.Longitude, &solarEfficiencie.Latitude,
		&solarEfficiencie.Annual, &solarEfficiencie.January,
		&solarEfficiencie.February, &solarEfficiencie.March,
		&solarEfficiencie.April, &solarEfficiencie.May,
		&solarEfficiencie.June, &solarEfficiencie.July,
		&solarEfficiencie.August, &solarEfficiencie.September,
		&solarEfficiencie.October, &solarEfficiencie.November,
		&solarEfficiencie.December)
	if err != nil {
		return solarEfficiencie, err
	}

	return solarEfficiencie, nil
}
