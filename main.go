package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type User struct {
	Username  string
	Statistic int
}

func main() {

	users := map[string][]User{
		"Users": {},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, users)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(1 * time.Second)
		log.Print(r.Header.Get("HX-Request"))
		username := r.PostFormValue("username")
		stat, err := strconv.Atoi(r.PostFormValue("statistic"))
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println(username, "-", stat)
		tmpl := template.Must(template.ParseFiles("index.html"))
		temp_user := User{Username: username, Statistic: stat}
		tmpl.ExecuteTemplate(w, "users-list-element", temp_user)
		users["Users"] = append(users["Users"], temp_user)
		fmt.Println(users["Users"])
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {
		sort.Slice(users["Users"], func(i, j int) bool {
			return users["Users"][i].Statistic > users["Users"][j].Statistic
		})
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "user-list", users)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-user/", h2)
	http.HandleFunc("/sort-users/", h3)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
