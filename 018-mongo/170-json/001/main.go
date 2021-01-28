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
	t = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	defer db.Close()
	r := httprouter.New()

	r.GET("/books", booksIndex)
	r.GET("/books/show/:isbn", getBook)
	r.GET("/books/create", booksCreateForm)
	r.GET("/books/update/:isbn", booksUpdateForm)

	r.POST("/books/create/process", booksCreateProcess)
	r.POST("/books/update/:isbn", booksUpdateProcess)
	r.POST("/books/delete/:isbn", booksDeleteProcess)

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

	err = t.ExecuteTemplate(w, "books.gohtml", bs)
	if err != nil {
		http.Error(w, http.StatusText(500)+" "+err.Error(), http.StatusInternalServerError)
		return
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
	t.ExecuteTemplate(w, "create.gohtml", nil)
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

func booksUpdateForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	b := Book{}
	isbn := p.ByName("isbn")

	b.Isbn = isbn
	row := db.QueryRow("select author,title,price from books where isbn = $1", b.Isbn)
	var ps string
	err := row.Scan(&b.Author, &b.Title, &ps)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	f64, err := strconv.ParseFloat(ps, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+": enter numbers for price", http.StatusNotAcceptable)
	}

	b.Price = float32(f64)

	data := struct {
		B Book
		I string
	}{
		B: b,
		I: isbn,
	}

	err = t.ExecuteTemplate(w, "update.gohtml", data)
	if err != nil {
		http.Error(w, http.StatusText(500)+" "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func booksUpdateProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	b := Book{}
	isbn := p.ByName("isbn")
	b.Isbn = r.FormValue("isbn")
	b.Author = r.FormValue("author")
	b.Title = r.FormValue("title")

	ps := r.FormValue("price")

	if b.Isbn == "" || b.Author == "" || b.Title == "" || ps == "" {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	f64, err := strconv.ParseFloat(ps, 32)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	b.Price = float32(f64)

	_, err = db.Exec("update books set isbn = $1, title = $2, author = $3, price = $4 where isbn = $5", b.Isbn, b.Title, b.Author, b.Price, isbn)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func booksDeleteProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName("isbn")

	db.Exec("delete from books where isbn = $1", isbn)

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
