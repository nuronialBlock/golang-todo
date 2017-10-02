package main

import (
	"golang-todo/data"
	"log"
)

const (
	URL = "mongodb://localhost:27017/todo-db"
)

func main() {
	err := data.OpenDBSession(URL)
	if err != nil {
		println("Couldn't connect the DB")
	}
	println("Got the session")
	todo := &data.Todo{}
	todo.Text = "Learn react-redux"
	err = todo.Insert()
	if err != nil {
		log.Fatalln(err)
	}
	todos, err := data.ListTodos()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(todos)
}
