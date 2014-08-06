package models;


import (
	
	"time"
	"io/ioutil"
	"net/http"
	"gopkg.in/mgo.v2"
	b "gopkg.in/mgo.v2/bson"
)


type Site struct {

	Id    b.ObjectId `bson:"_id"`

	Name	string `bson:"Name"`
	Url		string `bson:"Url"`

	CheckPoint string `bson:"CheckPoint"`
	Method	string `bson:"Method"`

	Email string `bson:"Email"`

	Duration	int `bson:"Duration"`
	Expiration int64 `bson:"Expiration"`

	Status    int `bson:"Status"`
	Disabled  bool `bson:"Disabled"`
}

func (s Site) HexId() string {

    return s.Id.Hex()
}

func (s Site) TStatus() (ss string) {

	switch(s.Status) {
		default: ss = "正常"
		case 0: ss = "不可用"
	}

	return
}

func (s Site) TExpiration() string {

	return time.Unix(s.Expiration, 0).Format("2006-01-02 15:04:05")
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

func ListSite() (result []Site) {

	_C().Find(nil).Iter().All(&result)
	return 
}

func DoSiteCheck() {

	for _, s := range ListSite() {

		now := time.Now().Unix()

		if s.Disabled || s.CheckPoint == "" || s.Duration <= 0 || s.Expiration > now {
			continue
		}


		s.Expiration = (now+int64(s.Duration))

		go (func () {

			u := s.CheckPoint
			m := s.Method
		
			var resp *http.Response
			var err error
			
			s.Status = 1

			if m == "GET" {
				resp,err = http.Get(u)
			}else if m == "POST" {
				resp,err = http.Post(u,"application/form-data-url", nil)
			}

			if err != nil {

				s.Status = 0

			}else{

				defer resp.Body.Close()
				if body,err1 := ioutil.ReadAll(resp.Body) ; err1 != nil || len(body) == 0 {
					s.Status = 0
				}
			}

			if s.Status == 0 && s.Email != "" {

				go NewEmail(s.Email, s.Name + "访问异常" , "请注意！")
			}

			UpdateSite(s)
			
		})()
	}
}

