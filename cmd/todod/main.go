package main

import (
	"golang-todo/data"
	"golang-todo/ui"
	"log"
	"net/http"
)

const (
	URL = "mongodb://localhost:27017/todo-db"
)

func main() {
	err := data.OpenDBSession(URL)
	if err != nil {
		println("Couldn't connect the DB")
	}

	http.Handle("/", ui.NewServer())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
