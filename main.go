package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo")
	})

	mux.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		fmt.Fprintf(w, "you chose URL "+url)
	})

	urlMap := map[string]string{
		"/c": "http://www.google.com",
		"/d": "http://www.gophercises.com",
	}

	handler := mapHandler(urlMap, mux)
	http.ListenAndServe(":8080", handler)
}

func mapHandler(urls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		url := req.URL.Path
		val, present := urls[url]

		if present {
			http.Redirect(resp, req, val, http.StatusTemporaryRedirect)
			return
		}

		fallback.ServeHTTP(resp, req)
	}
}
