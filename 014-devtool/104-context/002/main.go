package main

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type key string

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/foo", foo)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	log.Println(ctx)

	ctx = context.WithValue(ctx, key("uID"), 77)
	ctx = context.WithValue(ctx, key("fname"), "J")
	res := dbAccess(ctx)

	log.Println(res)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value(key("uID")).(int)
	return uid
}

func foo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	log.Println(ctx)
}
