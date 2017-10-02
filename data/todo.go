package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID        bson.ObjectId `bson:"_id"`
	Text      string        `bson:"text"`
	Completed bool          `bson:"completed"`

	CreatedAt  time.Time `bson:"created_at"`
	ModifiedAt time.Time `bson:"modified_at"`
}

func (t *Todo) Insert() error {
	t.ModifiedAt = time.Now()
	if t.ID == "" {
		t.ID = bson.NewObjectId()
		t.CreatedAt = t.ModifiedAt
	}
	_, err := sess.DB("").C("todos").UpsertId(t.ID, t)
	return err
}

func ListTodos() ([]Todo, error) {
	todos := []Todo{}
	err := sess.DB("").C("todos").Find(nil).All(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
