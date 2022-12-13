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
	mot        string
	nouveaumot string
	essais     int
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
		data := Jeu{
			mot:    hangman.MotAleatoire(),
			essais: 10,
		}
		tmpl2.Execute(w, data)
		var lettre string
		nouveaumot := hangman.MasquerMot(data.mot)
		for 0 != 1 {
			if hangman.MotEstTrouve(nouveaumot) {

			} else if data.essais == 0 {

			}
			if r.Method == http.MethodPost {
				if r.FormValue("lettre") != "" {
					lettre = r.FormValue("lettre")
				}
			}
			hangman.LancementDuJeu(data.essais, lettre, nouveaumot, data.mot)
			data.essais, lettre, nouveaumot = hangman.LancementDuJeu(data.essais, lettre, nouveaumot, data.mot)
		}
	})
	http.ListenAndServe(":80", nil)
}
