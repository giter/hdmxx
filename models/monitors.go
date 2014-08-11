package models;


import (
	"fmt"
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
	Count	int	`bson:"Count"`
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

	_C().Find(nil).Sort("_id").Iter().All(&result)
	return 
}

func DoSiteCheck() {

	for _, s := range ListSite() {

		now := time.Now().Unix()


		if s.Disabled || s.CheckPoint == "" || s.Duration <= 0 || s.Expiration > now {

			continue
		}


		s.Expiration = (now+int64(s.Duration))
		s.Count++
		UpdateSite(s)

		go (func (s Site) {

			u := s.CheckPoint
			m := s.Method
		
			var resp *http.Response
			var err error
			
			s.Status = 1

			fmt.Println("Processing " + u + "...")

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

				go NewEmail(s.Email, "网站访问异常" , "名称："+s.Name+"\r\n网址："+s.Url+"\r\n\r\n"+"请注意处理！")
			}

			UpdateSite(s)
			
		})(s)
	}
}

