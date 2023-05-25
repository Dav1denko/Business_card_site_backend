package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func allProject(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/allProject.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "AllProject", nil)
}

func loadStaticData() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
}

func handleFunc() {
	loadStaticData()
	http.HandleFunc("/", index)
	http.HandleFunc("/allProject/", allProject)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
