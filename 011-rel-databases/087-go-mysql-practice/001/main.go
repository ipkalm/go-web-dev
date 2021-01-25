package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "golang:golangpasswd@tcp(127.0.0.1:3308)/golang?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	defer db.Close()

	r := httprouter.New()
	r.GET("/", index)
	r.GET("/genres", genres)
	r.GET("/create", create)
	r.GET("/insert", insert)
	r.GET("/read", read)
	r.GET("/update", update)
	r.GET("/delete", del)
	r.GET("/drop", drop)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := io.WriteString(w, "index page\n")
	if err != nil {
		log.Panicln(err)
	}
}

func genres(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query(`select name from metal`)
	if err != nil {
		log.Panicln(err)
	}

	var answer, tmp string
	for rows.Next() {
		err = rows.Scan(&tmp)
		if err != nil {
			log.Panicln(err)
		}
		answer += tmp + "\n"
	}

	fmt.Fprintln(w, answer)
}

func create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stmt, err := db.Prepare(`create table bands(
		bID int not null auto_increment,
		name varchar(100) not null,
		primary key(bID))`)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		log.Panicln(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprintln(w, "created new table bands", n)
}

func insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stmt, err := db.Prepare(`insert into bands(name) values('Om')`)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		log.Panicln(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprintln(w, "insert rows: ", n)
}

func read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query(`select * from bands`)
	if err != nil {
		log.Panicln(err)
	}

	var n, tmp string
	for rows.Next() {
		err = rows.Scan(&n, &tmp)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Fprintln(w, n, tmp)
	}
}

func update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stmt, err := db.Prepare(`update bands set name='Sleep' where bID=1`)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		log.Panicln(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprintln(w, "update rows: ", n)
}

func del(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stmt, err := db.Prepare(`delete from bands where bID=1`)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		log.Panicln(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprintln(w, "delete rows: ", n)
}

func drop(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stmt, err := db.Prepare(`drop table bands`)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		log.Panicln(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprintln(w, "drop table: ", n)
}
