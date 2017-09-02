

package main

import (
  "net/http"
  "html/template"
)

const (
  PORT = ":8080"
)

func viewHandler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("./frontend/html/main.html")
	t.Execute(w,r)
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(PORT, nil)
}
