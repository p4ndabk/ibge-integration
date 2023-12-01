package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/p4ndabk/ibge-integration/app/ibge"
	"github.com/p4ndabk/ibge-integration/infra/database"
)

func CreateSolarEfficiencyTable() (bool, error) {
	fmt.Println("Creating solar_efficiencies table...")
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	script := `
		CREATE TABLE IF NOT EXISTS solar_efficiencies (
			id INT PRIMARY KEY,
			country VARCHAR(255),
			longitude DECIMAL(10, 6),
			latitude DECIMAL(10, 6),
			annual INT,
			january INT,
			february INT,
			march INT,
			april INT,
			may INT,
			june INT,
			july INT,
			august INT,
			september INT,
			october INT,
			november INT,
			december INT
		);
	`
	_, err = db.Exec(script)
	if err != nil {
		log.Printf("%q: %s\n", err, script)
		return false, err
	}

	return true, nil
}

func ImportSolarEfficiency() (bool, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	solarEfficiencies, err := getSolarEfficiencies()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	i := 0
	for _, solar := range solarEfficiencies.SolarEfficiencies{

		stmt, err := tx.Prepare("insert into solar_efficiencies(id, country, longitude, latitude, annual, january, february, march, april, may, june, july, august, september, october, november, december) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return false, err
		}
		_, err = stmt.Exec(solar.ID, solar.Country, solar.Longitude, solar.Latitude, solar.Annual, solar.January, solar.February, solar.March, solar.April, solar.May, solar.June, solar.July, solar.August, solar.September, solar.October, solar.November, solar.December)
		if err != nil {
			return false, err
		}
		i++
	}

	fmt.Println("Total of Solar efficiency imported: ", i)

	tx.Commit()
	database.CloseDB(db)
	return true, nil
}


func getSolarEfficiencies() (ibge.SolarEfficiencies, error) {
	var solarEfficiencies ibge.SolarEfficiencies

	jsonData, err := os.Open("storage/solar_efficiencies.json")
	if err != nil {
		return solarEfficiencies, err
	}

	decoder := json.NewDecoder(jsonData)
	if err := decoder.Decode(&solarEfficiencies); err != nil {
		return solarEfficiencies, err
	}

	return solarEfficiencies, nil
}