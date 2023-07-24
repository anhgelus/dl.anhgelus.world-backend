package main

import (
	"github.com/anhgelus/dl.anhgelus.world-backend/src"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).
		PathPrefix("/").
		HandlerFunc(src.Handle)
	r.PathPrefix("/").
		HandlerFunc(src.HandleNotAllowed)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
