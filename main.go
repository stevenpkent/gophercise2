package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	urls := map[string]string{
		"/c": "http://www.google.com",
		"/d": "http://www.gophercises.com",
	}

	jsonData := []byte(`[
		{ "path": "/e", "url": "https://www.msn.com/en-us" },
		{ "path": "/f", "url": "https://gobyexample.com/" }
	]`)

	var pathsToUrls []shortLink
	err := json.Unmarshal(jsonData, &pathsToUrls)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, elem := range pathsToUrls {
		urls[elem.Path] = elem.URL
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yo")
	})

	mux.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		fmt.Fprintf(w, "you chose URL "+url)
	})

	handler := mapHandler(urls, mux)
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

type shortLink struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
