package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ipkalm/go-web-dev/015-go-mongo/113-packages/002/model"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang=en>
	<head>
	<meta charset="UTF-8">
	<title>index</title>
	</head>
	<body>
	<a href="/user/987">go to: http://localhost:8080/user/987</a>
	</body>
	</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := model.User{
		Name:   "j",
		Gender: "m",
		Age:    32,
		ID:     p.ByName("id"),
	}
	ujs, err := json.Marshal(u)
	if err != nil {
		log.Panicln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", ujs)
}
