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
	Essais  int
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
	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl2.Execute(w, nil)
			return
		}
		var lettre string
		mot := hangman.MotAleatoire()
		nouveaumot := hangman.MasquerMot(mot)
		essais := 10
		for 0 != 1 {
			if hangman.MotEstTrouve(nouveaumot) {

			} else if essais == 0 {

			}
			if r.Method == http.MethodPost {
				if r.FormValue("lettre") != "" {
					lettre = r.FormValue("lettre")
				}
			}
			hangman.LancementDuJeu(essais, lettre, nouveaumot, mot)
			essais, lettre, nouveaumot = hangman.LancementDuJeu(essais, lettre, nouveaumot, mot)
		}
	})
	http.ListenAndServe(":80", nil)
}
