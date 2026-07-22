package main

import (
	"net/http"

	"Kitchen/handler"
	"Kitchen/store"
)

func main() {
	h := &handler.Handler{Store: store.NewMemoryStore()}
	if err := http.ListenAndServe(":8080", handler.WithCORS(h.Routes())); err != nil {
		panic(err)
	}
}
