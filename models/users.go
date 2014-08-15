package models;


import (
	"gopkg.in/mgo.v2"
	b "gopkg.in/mgo.v2/bson"
)

type User struct {

	Id string	`bson:"_id"`

	Account string `bson:"Account"`
	Password string `bson:"Password"`

	Email string	`bson:"Email"`
	Mobile string `bson:"Mobile"`
}

const COLL_USER = "user"

func UserColl() *mgo.Collection {
	return DB().C(COLL_USER)
}

func NewUser(u User) (err error) {
		err = UserColl().Insert(&u)
		return
}

func UpdateUser(u User) (err error) {

		_,err = UserColl().Upsert(b.M{"_id": u.Id},u)
		return
}

func ListUser()(r []User,err error) {

	err = UserColl().Find(nil).Sort("_id").Iter().All(&r)
	return
}

func UserLogin(Account string, Password string) (u User) {

	FindUser(b.M{"Account":Account,"Password":Password}).One(&u)
	return
}

func FindUser(query interface{}) *mgo.Query {
	return UserColl().Find(query)
}
