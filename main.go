package main

import (
	"github.com/anhgelus/dl.anhgelus.world-backend/src"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := os.Mkdir("/data", 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	r := mux.NewRouter()
	r.Methods(http.MethodGet).
		PathPrefix("/").
		HandlerFunc(src.Handle)
	r.PathPrefix("/").
		HandlerFunc(src.HandleNotAllowed)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Default().Println("Starting...")
	log.Fatal(srv.ListenAndServe())
}
