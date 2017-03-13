package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("postgres", "user=username dbname=basename sslmode=verify-full")
	//db, err := sql.Open("postgres", "postgres://username:password@localhost/basename?sslmode=verify-full")
	db, err := sql.Open("postgres", "postgres://cesar@localhost/cesar?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id,date_time,log_text FROM log")

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("rows",rows)

	for rows.Next() {
		var id int
		var dateTime time.Time
		var logText string

		err = rows.Scan(&id, &dateTime, &logText)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strconv.Itoa(id), dateTime, logText)
	}

	db.Close()

}
