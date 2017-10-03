package ui

import (
	"golang-todo/data"
	"html/template"
	"net/http"
)

func serveTodoList(w http.ResponseWriter, r *http.Request) {
	todos, err := data.ListTodos()
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	t := template.Must(template.ParseFiles("ui/listTodos.gohtml"))
	err = t.Execute(w, todos)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func init() {
	router.
		Methods(get).
		Path("/todos").
		HandlerFunc(serveTodoList)
}
