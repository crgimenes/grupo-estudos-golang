package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {

	source := "postgres://cesar@localhost/cesar?sslmode=disable"

	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", source)
	if err != nil {
		return
	}

	defer db.Close()

	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		return
	}

	q := ""

	var stmt *sql.Stmt

	stmt, err = tx.Prepare(q)
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	par := make([]interface{}, 0)

	_, err = stmt.Exec(par)
	if err != nil {
		return
	}

}
