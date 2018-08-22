package models

import (
	"errors"

	"github.com/chinahdkj/xorm"
)

type User struct {
	Id int64 `bson:"_id" xorm:"pk autoincr"`

	Account  string `bson:"Account index"`
	Password string `bson:"Password"`

	Name   string `bson:"Name"`
	Email  string `bson:"Email"`
	Mobile string `bson:"Mobile"`

	Admin bool `bson:"Admin"`
}

func (u User) UserName() (n string) {

	n = u.Account

	if len(u.Name) > 0 {
		n = u.Name
	}

	return
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

const COLL_USER = "user"

func (this *UserService) Table() *xorm.Session {
	return DB().Table(new(User))
}

func (this *UserService) Insert(u User) error {
	_, err := this.Table().Insert(u)
	return err
}

func (this *UserService) Update(u User) error {
	_, err := this.Table().Id(u.Id).Update(u)
	return err
}

func (this *UserService) List() (s []User, err error) {
	err = this.Table().Find(&s)
	return
}

func (this *UserService) Login(Account string, Password string) (u *User, err error) {

	var s []User = []User{}

	err = this.Table().And("account = ?", Account).And("password = ?", Password).Limit(2).Find(&s)

	if err != nil {
		return nil, err
	}

	if len(s) > 1 || len(s) == 0 {
		return nil, errors.New("登录失败!")
	}

	u = &s[0]
	return u, err
}
