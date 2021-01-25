package main

import (
	"context"
	"log"
	"net/http"
	"time"

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
	res, err := dbAccess(ctx)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		uid := ctx.Value(key("uID")).(int)
		time.Sleep(10 * time.Second)
		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func foo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	log.Println(ctx)
}
