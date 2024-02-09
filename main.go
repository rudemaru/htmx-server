package main

import (
	"html/template"
	"log"
	"net/http"
)

type Contact struct {
	Name  string
	Email string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		contacts := map[string][]Contact{
			"Contacts": {
				{Name: "LeBron James", Email: "lebronthegoat@gmail.com"},
				{Name: "Zaraki Kenpachi", Email: "thestrongest11@mail.ru"},
				{Name: "Jeff Bezos", Email: "mnogobabok@gmail.com"},
			},
		}
		tmpl.Execute(w, contacts)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(1 * time.Second)
		cname := r.PostFormValue("c-name")
		cemail := r.PostFormValue("c-email")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "contact-list-element", Contact{Name: cname, Email: cemail})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-contact/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
