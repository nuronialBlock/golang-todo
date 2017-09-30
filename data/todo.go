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
	// log.Println(*t)
	_, err := sess.DB("").C("texts").UpsertId(t.ID, t)
	return err
}
