package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func open(dbsource string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", dbsource)
	if err != nil {
		err = fmt.Errorf("error open db: %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("error ping db: %v", err)
	}
	return
}

func main() {
	/*************************
	 ** Abre banco de dados **
	 *************************/

	//dbsource := "postgres://postgres:password@example.com/testdb?sslmode=verify-full"
	dbsource := "postgres://postgres@localhost/testdb?sslmode=disable"
	db, err := open(dbsource)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db.DriverName())

	/************
	 ** Insert **
	 ************/

	/*
	   Na maioria dos casos você pode usar tanto
	   Exec como Query, a diferença esta no retorno,
	   Query é mais adequado para quando você quer
	   ler linhas de retorno.
	*/

	// Insert simples
	// -=-=-=-=-=-=-=

	sql := `INSERT INTO "clients" ("name","address") VALUES ($1,$2)`

	_, err = db.Exec(sql,
		"Tyrell Corporation",
		"TC Earth Headquarters")
	if err != nil {
		fmt.Println(err)
		return
	}

	// named insert
	// -=-=-=-=-=-=
	type client struct {
		Name    string `json:"name" db:"name"`
		Address string `json:"address" db:"address"`
	}

	namedSQL := `INSERT INTO "clients" ("name","address") VALUES (:name,:address)`

	_, err = db.NamedExec(namedSQL,
		client{
			Name:    "Cyberdyne Systems",
			Address: "2144 Kramer St",
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	// named insert e retorna ID
	// -=-=-=-=-=-=-=-=-=-=-=-=-
	id := 0
	rows, err := db.NamedQuery(`INSERT INTO "clients" ("name","address") VALUES (:name,:address) RETURNING id`,
		client{
			Name:    "Umbrella Corporation",
			Address: "545 S Birdneck RD STE 202B Virginia Beach, VA 23451",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("id", id)
	}
	err = rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert com transação
	// -=-=-=-=-=-=-=-=-=-=
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = tx.Exec(sql,
		"OCP Omni Consumer Products",
		"Delta City (formerly Detroit)")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = tx.Exec(sql,
		"Weyland-Yutani Corporation",
		"Weyland-Yutani Corporation HQ, Tokyo")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = tx.Exec(sql,
		"GeneCo",
		"401 N. Boonville Avenue Springfield")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert usando prepare
	// -=-=-=-=-=-=-=-=-=-=-
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = stmt.Exec("Black Mesa", "Black Mesa, New Mexico, USA")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = stmt.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// MustExec (se der erro quebra tudo)
	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

	db.MustExec(sql,
		"League of Industrial Nations",
		"CON-AM 27, Io, Jupter")

	db.MustExec(sql,
		"Aperture Laboratories",
		"Upper Michigan, USA")

	/************
	 ** Select **
	 ************/

	sql = `select "name", "address" from "clients" order by name`

	// Select simples
	// -=-=-=-=-=-=-=
	r, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	// list := []client{}
	for r.Next() {
		c := client{}                     // nova instancia para conter o cliente
		err = r.Scan(&c.Name, &c.Address) // popula nova instancia
		if err != nil {                   // verifica se erro
			fmt.Println(err)
			return
		}
		fmt.Println("Nome....:", c.Name)
		fmt.Println("Endereço:", c.Address)
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		// list = append(list, c)
	}
	err = r.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Select lendo item a item com StructScan
	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
	rows, err = db.Queryx(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	// list := []client{}
	for rows.Next() {
		c := client{}             // nova instancia para conter o cliente
		err = rows.StructScan(&c) // popula nova instancia
		if err != nil {           // verifica se erro
			fmt.Println(err)
			return
		}
		fmt.Println("Nome....:", c.Name)
		fmt.Println("Endereço:", c.Address)
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		// list = append(list, c)
	}
	err = rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Select lendo todos os itens de uma vez usando db.Select
	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

	rows, err = db.Queryx(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	list := []client{}
	err = db.Select(&list, sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range list {
		fmt.Println("Registro:", k+1) // não é o id :D
		fmt.Println("Nome....:", v.Name)
		fmt.Println("Endereço:", v.Address)
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	}

	// Select lendo apenas um item
	// -=-=-=-=-=-=-=-=-=-=-=-=-=-

	// limit 1
	sql = `select "name", "address" from "clients" limit 1`
	c := client{}
	err = db.Get(&c, sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Nome....:", c.Name)
	fmt.Println("Endereço:", c.Address)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	// count
	sql = `select count(*) from "clients"`
	count := 0
	err = db.Get(&count, sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("count:", count)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	/***********
	 ** Close **
	 ***********/

	err = db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
