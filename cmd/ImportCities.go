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

	script := `
		CREATE TABLE IF NOT EXISTS cities (
			code_ibge INTEGER,
			code_uf INTEGER,
			name TEXT,
			latitude REAL,
			longitude REAL,
			capital INTEGER
		);
	`
	_, err = db.Exec(script)
	if err != nil {
		log.Printf("%q: %s\n", err, script)
		return false, err
	}

	return true, nil
}

func ImportCities() {
	fmt.Println("Importing cities...")
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	cities, err := importCitiesFile()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	i := 0
	for _, city := range cities.Cities {

		stmt, err := tx.Prepare("insert into cities(code_ibge, code_uf, name, latitude, longitude, capital) values(?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(city.CodeIBGE, city.CodeUF, city.Name, city.Latitude, city.Longitude, city.Capital)
		if err != nil {
			panic(err)
		}
		i++
	}
	tx.Commit()
	database.CloseDB(db)
	fmt.Println("Total of cities imported: ", i)
	fmt.Println("Cities imported!")
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
