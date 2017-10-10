package ui

import (
	"fmt"
	"golang-todo/data"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2/bson"
)

type Key string

const (
	get     string = "GET"
	post    string = "POST"
	delete  string = "DELETE"
	sessKey Key    = "accountID"
)

var store = sessions.NewCookieStore(
	[]byte(securecookie.GenerateRandomKey(64)),
	[]byte(securecookie.GenerateRandomKey(32)))

var router = mux.NewRouter()

type Server struct {
	router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, "s")
	if err != nil {
		fmt.Println("herererere")
		fmt.Println(err)
		return
	}
	accID, ok := sess.Values["accountID"].(string)
	fmt.Println(accID)
	if ok {
		fmt.Println("here at check")
		acc, err := data.GetAccount(bson.ObjectIdHex(accID))
		if err != nil {
			log.Fatalln(err)
		}
		context.Set(r, sessKey, acc.ID.Hex())
		fmt.Println("Context: ", context.Get(r, sessKey))
		context.Clear(r)
		// context.Set(r, "accID", acc.ID.Hex())
	} else {
		// key := "accountID"
		sess.Values["accountID"] = ""
		sess.Save(r, w)
		context.Clear(r)
	}
	fmt.Println("Going for mux")
	s.router.ServeHTTP(w, r)
}

func NewServer() *Server {
	return &Server{
		router,
	}
}
