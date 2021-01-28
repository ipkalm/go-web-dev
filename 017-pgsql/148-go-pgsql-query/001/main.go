package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://bond:bond123@localhost:5433/company?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	rows, err := db.Query("select * from books")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i, t, a, p string
		err = rows.Scan(&i, &t, &a, &p)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("***\nisbn:\t%v\ntitle:\t%v\nauthor:\t%v\nprice:\t%v\n***\n\n", i, t, a, p)
	}

	log.Println("conected to pg")
}
