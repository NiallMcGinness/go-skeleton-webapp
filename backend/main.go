package main

import (
	"html/template"
	"net/http"
)

const (
	PORT = ":9000"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../views/main.html")
	t.Execute(w, r)
}

func main() {
	// serve static css and js files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(PORT, nil)
}
