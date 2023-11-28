package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/p4ndabk/ibge-integration/app/ibge"
	"github.com/p4ndabk/ibge-integration/infra/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
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
		return
	}
	cities, err := ibge.AllCiteis()
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	for _, city := range cities.Cities {
		stmt, err := tx.Prepare("insert into cities(code_ibge, code_uf, name, latitude, longitude, capital) values(?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}	
		defer stmt.Close()
		_, err = stmt.Exec(city.CodeIBGE, city.CodeUF, city.Name, city.Latitude, city.Longitude, city.Capital)
		if err != nil {
			panic(err)
		}
	}
	tx.Commit()
}
