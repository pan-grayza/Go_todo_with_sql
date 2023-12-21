package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Body  string
}

func main() {
	fmt.Println("Go app...")

	loadTodos := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./cmd/web/templates/index.html"))
		todos := map[string][]Todo{
			"Todos": {},
		}
		tmpl.Execute(w, todos)
	}

	addTodo := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		tmpl := template.Must(template.ParseFiles("./cmd/web/templates/index.html"))
		tmpl.ExecuteTemplate(w, "todo-list-element", Todo{Title: title, Body: body})
	}

	deleteTodo := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		tmpl := template.Must(template.ParseFiles("./cmd/web/templates/index.html"))
		tmpl.ExecuteTemplate(w, "todo-list-element", Todo{Title: title, Body: body})
	}

	// define handlers
	http.HandleFunc("/", loadTodos)
	http.HandleFunc("/add-todo/", addTodo)
	http.HandleFunc("/delete-todo/", deleteTodo)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
