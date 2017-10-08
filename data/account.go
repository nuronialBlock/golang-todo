package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID       bson.ObjectId   `bson:"_id"`
	Username string          `bson:"username"`
	Password AccountPassword `bson:"password"`

	CreatedAt  time.Time `bson:"created_at"`
	ModifiedAt time.Time `bson:"modified_at"`
}
