package main

import (
	"html/template"
	"net/http"
)

type Car struct {
	Brand string
	Model string
	Power int
	Available bool
}

type Garage struct {
	Name string
	Cars []Car
}

func main() {

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Garage{
			Name: "Alan's Cars",
			Cars: []Car{
				{Brand: "Audi", Model: "TT", Power: 245 , Available: false},
				{Brand: "Lamborghini", Model: "Aventador SVJ",Power: 770 , Available: true},
				{Brand: "Ferrari", Model: "F8 Spider",Power: 720 , Available: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
	
}