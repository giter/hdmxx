package models;


import (

	"time"
	"io/ioutil"
	"net/http"
	"gopkg.in/mgo.v2"
	b "gopkg.in/mgo.v2/bson"
)


type Site struct{

	Id    b.ObjectId `bson:"_id"`

	Name	string `bson:"Name"`
	Url		string `bson:"Url"`
	Method	string `bson:"Method"`

	Email string `bson:"Email"`

	Duration	int `bson:"Duration"`
	Expiration int64 `bson:"Expiration"`
}

func (s Site) HexId() string {

    return s.Id.Hex()
}

const COLL_SITE = "site"

func _C() *mgo.Collection {

	return DB().C(COLL_SITE)
}

func NewSite(s Site){
	
	_C().Insert(s)
}

func UpdateSite(s Site){

	_C().Upsert(b.M{"_id": s.Id},s)
}

func UpdateSiteExpiration(Id b.ObjectId, Expiration int64) {

	_C().Update(b.M{"_id": Id},b.M{"$set": b.M{"Expiration": Expiration}})
}

func ListSite() (result []Site) {

	_C().Find(nil).Iter().All(&result)
	return 
}

func DoSiteCheck() {

	for _, s := range ListSite() {

		now := time.Now().Unix()

		if s.Duration <= 0 || s.Expiration > now || s.Email == "" {
			continue
		}

		UpdateSiteExpiration(s.Id,(now+int64(s.Duration)))

		go (func () {

			m := s.Method
			u := s.Url
		
			var resp *http.Response
			var err error

			if m == "GET" {
				resp,err = http.Get(u)
			}else if m == "POST" {
				resp,err = http.Post(u,"application/form-data-url", nil)
			}

			if err != nil {

				NewEmail(s.Email, s.Name + "访问异常" , "请注意！")

			}else{

				defer resp.Body.Close()

				body,err1 := ioutil.ReadAll(resp.Body)

				if err1 != nil || len(body) == 0 {
					NewEmail(s.Email, s.Name + "访问异常" , "请注意！")
				}
			}
		})()
	}
}

