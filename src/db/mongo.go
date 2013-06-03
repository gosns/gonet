package db

import (
	"cfg"
)

import (
	"labix.org/v2/mgo"
)

var Mongo *mgo.Session

func init() {
	config := cfg.Get()
	session, err := mgo.Dial(config["mongo_host"])

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	Mongo = session
}

//------------------------------------------------ for very simple use
func Collection(name string) *mgo.Collection {
	config := cfg.Get()
	return Mongo.DB(config["mongo_db"]).C(name)
}
