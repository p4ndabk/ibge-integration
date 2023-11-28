package database

import "database/sql"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Query(query string) (*sql.Rows, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}