package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS companies (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		panic(err)
	}

	// evil megacorporations
	companies := []string{
		"Apperture Science",            // Portal
		"Cyberdyne Systems",            // Terminator
		"Multi-National United (MNU)",  // District 9
		"Omni Consumer Products (OCP)", // Robocop
		"Tyrell Corporation",           // Blade Runner
		"Umbrella Corporation",         // Resident Evil
		"Wallace Corporation",          // Blade Runner 2049
		"Weyland-Yutani Corporation",   // Alien
	}

	for _, company := range companies {
		_, err = db.Exec("INSERT INTO companies (name) VALUES (?)", company)
		if err != nil {
			panic(err)
		}
	}

	// Query the database
	rows, err := db.Query("SELECT id, name FROM companies")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		println("ID:", id, "Name:", name)
	}
}
