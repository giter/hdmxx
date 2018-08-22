package models

import "github.com/chinahdkj/xorm"
import _ "github.com/mattn/go-sqlite3"

var db *xorm.Engine = nil

func init() {

	d, err := xorm.NewEngine("sqlite3", "./jxbskj.db")

	if err != nil {
		panic(err)
	}

	db = d

	db.Sync2(new(User))
}

func DB() *xorm.Engine {
	return db
}
