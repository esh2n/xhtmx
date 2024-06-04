package main

import (
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	Id      int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{1, "Buy Milk"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		tmpl := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		data["Todos"] = append(data["Todos"], todo)

		tmpl.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
