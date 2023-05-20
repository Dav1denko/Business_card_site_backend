package main

import (
	"fmt"
	"net/http"
)

type User struct {
}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test GOGOGOGo")
}

func contatts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contatdds")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contatts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
