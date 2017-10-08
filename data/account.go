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

func GetAccountByUsername(username string) (Account, error) {
	acc := Account{}
	err := sess.DB("").C("accounts").Find(bson.M{"username": username}).One(&acc)
	if err != nil {
		return Account{}, err
	}
	return acc, nil
}

func (a *Account) Insert() error {
	a.ModifiedAt = time.Now()
	if a.ID == "" {
		a.CreatedAt = a.ModifiedAt
		a.ID = bson.NewObjectId()
	}
	_, err := sess.DB("").C("accounts").UpsertId(a.ID, a)
	return err
}
