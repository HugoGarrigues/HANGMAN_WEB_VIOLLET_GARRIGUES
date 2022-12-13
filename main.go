package main

import (
	"hangman-web/hangman"
	"html/template"
	"net/http"
)

type User struct {
	Pseudo  string
	Success bool
	Niveau  string
}

type Jeu struct {
	Mot        string
	NouveauMot string
	Essais     int
}

func main() {

	tmpl1 := template.Must(template.ParseFiles("index.html"))
	tmpl2 := template.Must(template.ParseFiles("idCard.html"))

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
		details := User{
			Pseudo:  r.FormValue("pseudo"),
			Niveau:  r.FormValue("niveau"),
			Success: true,
		}
		tmpl1.Execute(w, details)
	})
	data := Jeu{
		Mot:        hangman.MotAleatoire(),
		NouveauMot: hangman.MasquerMot(hangman.MotAleatoire()),
		Essais:     10,
	}
	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {

		if hangman.LettreEstPresente(r.FormValue("letter"), data.Mot) && len(r.FormValue("letter")) == 1 {
			data.Essais--
		}
		tmpl2.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
