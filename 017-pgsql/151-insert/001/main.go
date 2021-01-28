package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// Book is just a book
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

var (
	db *sql.DB
	t  *template.Template
)

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
	t = template.Must(template.ParseFiles("create.gohtml"))
}

func main() {
	defer db.Close()
	r := httprouter.New()

	r.GET("/books", booksIndex)
	r.GET("/books/show/:isbn", getBook)
	r.GET("/books/create", booksCreateForm)

	r.POST("/books/create/process", booksCreateProcess)

	log.Fatal(http.ListenAndServe(":8080", r))
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
		err = rows.Scan(&b.Isbn, &b.Title, &b.Author, &b.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bs = append(bs, b)
	}

	for _, b := range bs {
		fmt.Fprintf(w, "***\nisbn:\t%v\ntitle:\t%v\nauthor:\t%v\nprice:\t%.2f\n***\n\n", b.Isbn, b.Title, b.Author, b.Price)
	}
}

func getBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName("isbn")
	if isbn == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	row := db.QueryRow("select * from books where isbn = $1", isbn)

	b := Book{}
	err := row.Scan(&b.Isbn, &b.Title, &b.Author, &b.Price)
	switch err {
	case sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case nil:
		break
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "***\nisbn:\t%v\ntitle:\t%v\nauthor:\t%v\nprice:\t%.2f\n***\n\n", b.Isbn, b.Title, b.Author, b.Price)
}

func booksCreateForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t.Execute(w, nil)
}

func booksCreateProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b := Book{}
	b.Isbn = r.FormValue("isbn")
	b.Author = r.FormValue("author")
	b.Title = r.FormValue("title")
	tp := r.FormValue("price")

	if b.Author == "" || b.Title == "" || b.Isbn == "" || tp == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(tp, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"need price", http.StatusNotAcceptable)
		return
	}
	b.Price = float32(f64)

	res, err := db.Exec("insert into books (isbn,title,author,price) values ($1,$2,$3,$4)", b.Isbn, b.Title, b.Author, b.Price)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	_, err = res.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
