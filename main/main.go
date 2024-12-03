package main

import (
	"fmt"
	"net/http"
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
	fmt.Fprintf(page, "Главная страница")
}

func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Контакты")
}
