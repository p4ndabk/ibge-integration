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

func CreateCitiesTable() (bool, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	createCitiesTable := `
		CREATE TABLE IF NOT EXISTS cities (
			code_ibge INTEGER,
			code_uf INTEGER,
			name TEXT,
			latitude REAL,
			longitude REAL,
			capital INTEGER
		);
	`
	_, err = db.Exec(createCitiesTable)
	if err != nil {
		log.Printf("%q: %s\n", err, createCitiesTable)
		return false, err
	}

	return true, nil
}

func ImportCities() (bool, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	cities, err := importCitiesFile()
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
	for _, city := range cities.Cities {

		stmt, err := tx.Prepare("insert into cities(code_ibge, code_uf, name, latitude, longitude, capital) values(?, ?, ?, ?, ?, ?)")
		if err != nil {
			return false, err
		}
		_, err = stmt.Exec(city.CodeIBGE, city.CodeUF, city.Name, city.Latitude, city.Longitude, city.Capital)
		if err != nil {
			return false, err
		}
		i++
	}

	fmt.Println("Total of cities imported: ", i)

	tx.Commit()
	database.CloseDB(db)
	return true, nil
}

func importCitiesFile() (ibge.Cities, error) {
	var cities ibge.Cities

	jsonData, err := os.Open("storage/cities.json")
	if err != nil {
		return cities, err
	}

	decoder := json.NewDecoder(jsonData)
	if err := decoder.Decode(&cities); err != nil {
		return cities, err
	}

	return cities, nil
}
