package books

import (
	"log"
	"net/http"

	"github.com/ipkalm/go-web-dev/018-mongo/171-books-app-mongo/config"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	err = config.TPL.ExecuteTemplate(w, "books.gohtml", bks)
	if err != nil {
		log.Panic(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = config.TPL.ExecuteTemplate(w, "show.gohtml", bk)
	if err != nil {
		log.Panic(err)
	}
}

func Create(w http.ResponseWriter, _ *http.Request) {
	err := config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
	if err != nil {
		log.Panic(err)
	}
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	err = config.TPL.ExecuteTemplate(w, "created.gohtml", bk)
	if err != nil {
		log.Panic(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := OneBook(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = config.TPL.ExecuteTemplate(w, "update.gohtml", bk)
	if err != nil {
		log.Panic(err)
	}
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	err = config.TPL.ExecuteTemplate(w, "updated.gohtml", bk)
	if err != nil {
		log.Panic(err)
	}
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
