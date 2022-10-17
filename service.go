package main

import (
	"database/sql"
)

func getQeury(query string) *sql.Rows {
	db := database()
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	return rows
}

func makingBuy(query string) *sql.Rows {
	db := database()
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	return rows
}

func makingSell(query string) *sql.Rows {
	db := database()
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	return rows
}
