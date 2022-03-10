package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:666/")
	http.ListenAndServe(":666", nil)
}
