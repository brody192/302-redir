package main

import (
	"cmp"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "root path, nothing to see here")
	})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})

	mux.HandleFunc("GET /redirect", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/hello")

		w.WriteHeader(302)
	})

	port := cmp.Or(os.Getenv("PORT"), "3001")

	fmt.Println("starting server on port:", port)

	log.Fatal(http.ListenAndServe((":" + port), mux))
}
