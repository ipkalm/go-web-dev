package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "golang:golangpasswd@tcp(127.0.0.1:3308)/golang?charset=utf8")
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panicln(err)
	}

	r := httprouter.New()
	r.GET("/", index)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := io.WriteString(w, "success")
	if err != nil {
		log.Panicln(err)
	}
}
