package main

import (
	"fmt"
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
	Mot          string
	NouveauMot   string
	Essais       int
	Victoire     bool
	Defaite      bool
	Liste_Lettre string
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
		Mot:          hangman.MotAleatoire(),
		NouveauMot:   hangman.MasquerMot(hangman.MotAleatoire()),
		Essais:       0,
		Victoire:     false,
		Defaite:      false,
		Liste_Lettre: "",
	}

	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if niveau := r.FormValue("niveau"); niveau != "" {
				fmt.Println("bonjour")
				if niveau == "facile" {
					data.Mot = hangman.MotAleatoire()
					data.Mot = hangman.MettreEnMajuscule(data.Mot)
					data.NouveauMot = hangman.MasquerMot(data.Mot)
				} else if niveau == "moyen" {
					data.Mot = hangman.MotAleatoire1()
					data.Mot = hangman.MettreEnMajuscule(data.Mot)
					data.NouveauMot = hangman.MasquerMot(data.Mot)
				} else if niveau == "difficile" {
					data.Mot = hangman.MotAleatoire2()
					data.Mot = hangman.MettreEnMajuscule(data.Mot)
					data.NouveauMot = hangman.MasquerMot(data.Mot)
				}
			}
		}
		fmt.Println(r.FormValue("niveau"))
		var lettre string
		if data.Essais != 10 && data.Mot != data.NouveauMot {
			if r.Method == http.MethodPost {
				if r.FormValue("lettre") != " " {
					lettre = r.FormValue("lettre")
					lettre = hangman.MettreEnMajuscule(lettre)
					if hangman.LettreEstPresente(lettre, data.Mot) {
						data.NouveauMot = hangman.AfficheMotAvecLettreTrouvee(lettre, data.Mot, data.NouveauMot)
					} else {
						data.Liste_Lettre = hangman.Ajout_lettre(lettre, data.Liste_Lettre)
						data.Essais++
					}
				}
			}
			if data.Essais == 10 {
				data.Defaite = true
			}
			if data.Mot == data.NouveauMot {
				data.Victoire = true
			}
		}
		tmpl2.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
