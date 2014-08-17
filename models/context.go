package models;

import (
	"gopkg.in/mgo.v2"
)

const TYPE_HTTP = "HTTP"
const TYPE_TCP  = "TCP"
const TYPE_UDP  = "UDP"

const DB_NAME = "test"

var _S , _ = mgo.Dial("127.0.0.1")
var _Init = false

func InitM() {

	_Init = true
	_S.SetMode(mgo.Monotonic, true)
}

func M() *mgo.Session {

	if _S != nil {

		if !_Init {
			InitM()
		}

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
