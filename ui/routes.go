package ui

import (
	"fmt"
	"golang-todo/data"
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	mgo "gopkg.in/mgo.v2"
)

func serveTodoList(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, "s")
	fmt.Println("Serving for ID:", sess.Values["accountID"])

	fmt.Println("Context: ", context.Get(r, sessKey).(string))

	if err != nil {
		fmt.Println(err)
	}
	todos, err := data.ListTodos()
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Println(todos)
	t := template.Must(template.ParseFiles("ui/listTodos.gohtml"))
	err = t.Execute(w, todos)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	acc, err := data.GetAccountByUsername(username)
	if err == mgo.ErrNotFound {
		fmt.Println("User:", username, "doesn't exist. Need to sign up fist.")
		return
	}

	matched := acc.Password.Match(password)
	if !matched {
		fmt.Println("Not gonna happen Bro, this is highly encrypted :p")
		return
	}

	fmt.Println("Welcome,", username)
	session, err := store.Get(r, "s")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	context.Set(r, sessKey, acc.ID.Hex())
	context.Clear(r)
	fmt.Println("Context: ", context.Get(r, sessKey))
	session.Values["accountID"] = acc.ID.Hex()
	session.Save(r, w)
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	acc, err := data.GetAccountByUsername(username)
	if err != mgo.ErrNotFound {
		fmt.Println("User", username, "already exists in the system.")
		return
	}

	acc.Username = username
	acc.Password, err = data.NewAccountPassword(password)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = acc.Insert()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(acc)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, "s")
	if err != nil {
		fmt.Println(err)
	}
	sess.Values["accountID"] = ""
}

func init() {
	router.
		Methods(get).
		Path("/todos").
		HandlerFunc(serveTodoList)
	router.
		Methods(post).
		Path("/login").
		HandlerFunc(HandleLogin)
	router.
		Methods(post).
		Path("/logout").
		HandlerFunc(HandleLogout)
	router.
		Methods(post).
		Path("/signup").
		HandlerFunc(HandleSignup)
}
