package models;

import (
	"gopkg.in/mgo.v2"
)

const DB_NAME = "test"

var _S , _ = mgo.Dial("127.0.0.1")

func M() *mgo.Session {

	if _S != nil {
		return _S
	}

	panic("No Connection!")

}

var _D = M().DB(DB_NAME)

func DB() *mgo.Database {

	if _D != nil {
		return _D
	}

	panic("No Connection!")
}