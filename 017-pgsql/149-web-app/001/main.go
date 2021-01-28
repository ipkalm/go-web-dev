package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// Book is just a book
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://bond:bond123@localhost:5433/company?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	log.Println("connected to pgdb")
}

func main() {
	defer db.Close()
	r := httprouter.New()
	r.GET("/books", booksIndex)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func booksIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("select * from books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bs = []Book{}
	for rows.Next() {
		b := Book{}
		err = rows.Scan(&b.isbn, &b.title, &b.author, &b.price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bs = append(bs, b)
	}

	for _, b := range bs {
		fmt.Printf("***\nisbn:\t%v\ntitle:\t%v\nauthor:\t%v\nprice:\t%.2f\n***\n\n", b.isbn, b.title, b.author, b.price)
	}
}
