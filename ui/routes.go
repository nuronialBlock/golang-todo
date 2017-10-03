package ui

import (
	"fmt"
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
	fmt.Println("LALALA")
	err = t.Execute(w, todos)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// fmt.Println(todos)
}

func init() {
	router.
		Methods(get).
		Path("/todos").
		HandlerFunc(serveTodoList)
}
