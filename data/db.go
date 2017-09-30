package data

import "gopkg.in/mgo.v2"

var sess *mgo.Session

func OpenDBSession(url string) (err error) {
	sess, err = mgo.Dial(url)
	return err
}
