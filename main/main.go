package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
	"www/article"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	handleRequest()
}

type ErrorPageData struct {
	ErrorTitle   string
	ErrorMessage string
}

func handleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/contacts", contacts)
	http.HandleFunc("/about", about)
	http.HandleFunc("/save_article", save_article)

	http.ListenAndServe(":8080", nil)
}

func index(page http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	db, err := sql.Open("mysql", "root:Anama654!@tcp(127.0.0.1:3308)/golang")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("Select * From articles")

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var post article.Article

		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)

		if err != nil {
			panic(err)
		}
	}

	t.ExecuteTemplate(page, "index", nil)
}

func save_article(page http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	var errDate ErrorPageData

	if title == "" {
		errDate.ErrorTitle = "Ошибка заполнения формы"
		errDate.ErrorMessage += "[Заголовок не может быть пустым] "
	}
	if anons == "" {
		errDate.ErrorTitle = "Ошибка заполнения формы"
		errDate.ErrorMessage += " [Анонс не может быть пустым] "
	}
	if len(full_text) < 10 {
		errDate.ErrorTitle = "Ошибка заполнения формы"
		errDate.ErrorMessage += " [Текст статьи слишком короткий]"
	}

	if title != "" && anons != "" && len(full_text) >= 10 {
		db, err := sql.Open("mysql", "root:Anama654!@tcp(127.0.0.1:3308)/golang")

		if err != nil {
			panic(err)
		}

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (title, anons, full_text) VALUES ('%s', '%s', '%s')", title, anons, full_text))
		defer insert.Close()

		http.Redirect(page, r, "/", http.StatusSeeOther)
	}

	if errDate.ErrorTitle != "" {
		errorPage(page, r, errDate)
	}
}

func create(page http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	t.ExecuteTemplate(page, "create", nil)
}

func articles(page http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/articles.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	t.ExecuteTemplate(page, "articles", nil)
}

func about(page http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/about.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	t.ExecuteTemplate(page, "about", nil)
}

func errorPage(page http.ResponseWriter, r *http.Request, errDate ErrorPageData) {
	t, err := template.ParseFiles("templates/errorPage.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	data := errDate

	t.ExecuteTemplate(page, "errorPage", data)
}

func contacts(page http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/contacts.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	t.ExecuteTemplate(page, "contacts", nil)
}
