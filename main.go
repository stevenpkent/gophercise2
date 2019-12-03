package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo")
	})

	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		fmt.Fprintf(w, "you chose URL "+url)
	})

	http.ListenAndServe(":8080", mux)
}
