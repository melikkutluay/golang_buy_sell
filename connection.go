package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "keycloak"
	password = "password"
	dbname   = "metadata"
)

func database() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s "+"dbname=%s "+"sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	//defer db.Close()
	return db
}
