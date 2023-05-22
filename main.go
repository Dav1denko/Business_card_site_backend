package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/home_page.html")
	myName := "Lolsx"
	tmpl.Execute(w, myName)

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
