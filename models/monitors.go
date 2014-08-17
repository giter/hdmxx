package models;


import (
	"time"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	b "gopkg.in/mgo.v2/bson"
)


type Site struct {

	Id    b.ObjectId `bson:"_id"`

	Name	string `bson:"Name"`
	Url		string `bson:"Url"`

	Type string `bson:"Type"`

	//HTTP
	CheckPoint string `bson:"CheckPoint"`
	Method	string `bson:"Method"`

	//TCP/UDP
	Address string `bson:"Address"`
	Port int `bson:"Port"`
	Input string `bson:"Input"`
	Result string `bson:"Result"`

	Duration	int `bson:"Duration"`
	Expiration int64 `bson:"Expiration"`
	Run	int64	`bson:"Run"`

	Status    int `bson:"Status"`
	Disabled  bool `bson:"Disabled"`
	Count	int	`bson:"Count"`

	Users []User `bson:"Users"`
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

func (s Site) TRun() string {
	return time.Unix(s.Run, 0).Format("2006-01-02 15:04:05")
}

const COLL_SITE = "site"

func SiteColl() *mgo.Collection {

	return DB().C(COLL_SITE)
}

func NewSite(s Site) error{

	return SiteColl().Insert(s)
}

func UpdateSite(s Site) (err error){

	_,err = SiteColl().Upsert(b.M{"_id": s.Id},s)
	return
}

func FindSite(query interface{}) *mgo.Query {

	return SiteColl().Find(query);
}

func ListSite() (result []Site,err error) {

	err = SiteColl().Find(nil).Sort("_id").Iter().All(&result)
	return
}

func GetSite(Id string)(s *Site, err error){

	err = SiteColl().Find(b.M{"_id":b.ObjectIdHex(Id)}).One(&s)
	return
}

func DoSiteCheck() {

	sites, err := ListSite();
	if err != nil {
		return
	}

	for _, s := range sites {

		now := time.Now().Unix()


		if s.Disabled || s.Duration <= 0 || s.Expiration > now {

			continue
		}


		s.Expiration = (now+int64(s.Duration))
		s.Run = now
		s.Count++
		UpdateSite(s)

		go (func (s Site) {

			beego.Info("Processing " + s.Url + " ......")


			switch(s.Type){
			case "HTTP":
				s.Status, _ = CheckHttp(s.Method,s.CheckPoint);
			case "TCP":
				s.Status, _ = CheckNet("tcp",s.Address,s.Port,s.Input,s.Result)
			}

			UpdateSite(s)

			Title := "网站访问异常"
			Content := "名称："+s.Name+"\r\n网址："+s.Url+"\r\n\r\n"+"请注意处理！"

			beego.Info(s.Status)

			if s.Status == 0 && s.Users != nil && len(s.Users) > 0 {

				for _, u := range(s.Users) {

					if u.Email != "" {
						go NewEmail(u.Email, Title, Content)
					}
				}

			}

		})(s)
	}
}

