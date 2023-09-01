package main

import (
	"e1/database/bbolt"
	"fmt"
	"log"
)

func main() {

	db, err := bbolt.New("teste.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Update("jobs", "testKey", []byte("test 123"))
	if err != nil {
		log.Fatalf("Uodate(): %v\n", err)
	}

	b, err := db.View("jobs", "testKey")
	if err != nil {
		log.Fatalf("View(): %v\n", err)
	}
	fmt.Println("retorno:", string(b))

}
