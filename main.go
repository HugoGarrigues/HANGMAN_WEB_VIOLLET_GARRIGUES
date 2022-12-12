package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Pseudo  string
	Success bool
	Niveau  string
	Essais  int
}

func main() {

	tmpl1 := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
		details := User{
			Pseudo:  r.FormValue("pseudo"),
			Niveau:  r.FormValue("1 ou 2 ou 3"),
			Success: true,
		}
		tmpl1.Execute(w, details)
	})
	http.ListenAndServe(":80", nil)
}
