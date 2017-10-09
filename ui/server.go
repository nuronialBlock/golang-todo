package ui

import (
	"fmt"
	"golang-todo/data"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2/bson"
)

const (
	get    string = "GET"
	post   string = "POST"
	delete string = "DELETE"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

var router = mux.NewRouter()

type Server struct {
	router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, "S")
	fmt.Println("test context:", context.Get(r, "accID"))
	if err != nil {
		fmt.Println(err)
		return
	}
	accID, ok := sess.Values["accountID"].(string)
	fmt.Println(accID)
	if ok {
		fmt.Println("here")
		acc, err := data.GetAccount(bson.ObjectIdHex(accID))
		if err != nil {
			log.Fatalln(err)
		}
		context.Set(r, "accID", acc.ID.Hex())
		fmt.Println("test context:", context.Get(r, "accID"))
	}
	// else {
	// 	key := "accountID"
	// 	delete(sess.Values, key)
	// 	sess.Save(r, w)
	// 	context.Clear(r)
	// }
	s.router.ServeHTTP(w, r)
}

func NewServer() *Server {
	return &Server{
		router,
	}
}
