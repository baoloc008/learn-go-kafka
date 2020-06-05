package main

import (
	"log"
	"net/http"
)

func main() {

	r := http.NewServeMux()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	s := http.Server{
		Addr:":9620",
		Handler:r,
	}

	log.Fatal(s.ListenAndServe())
}
