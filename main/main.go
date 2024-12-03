package main

import (
	"fmt"
	"net/http"
	"www/myuser"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func home_page(page http.ResponseWriter, r *http.Request) {
	ivan := myuser.MyUser{
		Name:       "Иван",
		Age:        30,
		Money:      1000,
		Avg_grades: 4.2,
		Happiness:  75.0,
	}

	ivan.SetNewName("Мага")
	fmt.Fprintf(page, ivan.String())
}

func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Контакты")
}
