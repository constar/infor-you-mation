package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

type User struct {
	Id       int `orm:"pk;auto"`
	Username string
	Pwd      string
}

func (u *User) String() string {
	return fmt.Sprintf("Id:%d, Username:%s, Pwd:%s", u.Id, u.Username, u.Pwd)
}

func getOrm() beedb.Model {
	db, err := sql.Open("mysql", "inforyoumation@tcp(127.0.0.1:3306)/inforyoumation?charset=utf8")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func RegisterUser(user User) error {
	orm := getOrm()
	glog.Info("RegisterUser ", user)
	err := orm.Save(&user)
	return err
}

func ValidateUser(user User) error {
	orm := getOrm()
	var u User
	orm.Where("username=? and pwd=?", user.Username, user.Pwd).Find(&u)
	if u.Username == "" {
		return errors.New("用户名或者密码错误!")
	}
	return nil
}
