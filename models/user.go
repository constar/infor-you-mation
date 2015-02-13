package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoHost   = "127.0.0.1"
	mongoDBName = "inforyoumation"
	userColName = "user"
)

type User struct {
	Id       bson.ObjectId "_id"
	Username string
	Password string
}

func (u *User) String() string {
	return fmt.Sprintf("Id:%d, Username:%s, Password:%s", u.Id, u.Username, u.Password)
}

type ConnectOption struct {
	host   string
	dbSess *mgo.Session
	dbName string
}

var connOption ConnectOption

func Connect(host string, dbName string) error {
	var err error
	connOption.host = host
	connOption.dbSess, err = mgo.Dial(host)
	connOption.dbName = dbName
	return err
}

func init() {
	err := Connect(mongoHost, mongoDBName)
	if err != nil {
		panic(err)
	}
}

func RegisterUser(username string, passwd string) error {
	c := connOption.dbSess.DB(connOption.dbName).C(userColName)
	u := User{
		bson.NewObjectId(),
		username,
		passwd,
	}
	return c.Insert(&u)
}

func ValidateUser(user User) error {
	return nil
}

//func getOrm() beedb.Model {
//	db, err := sql.Open("mysql", "inforyoumation@tcp(127.0.0.1:3306)/inforyoumation?charset=utf8")
//	if err != nil {
//		panic(err)
//	}
//	orm := beedb.New(db)
//	return orm
//}
//
//func RegisterUser(user User) error {
//	orm := getOrm()
//	glog.Info("RegisterUser ", user)
//	err := orm.Save(&user)
//	return err
//}
//
//func ValidateUser(user User) error {
//	orm := getOrm()
//	var u User
//	orm.Where("username=? and pwd=?", user.Username, user.Password).Find(&u)
//	if u.Username == "" {
//		return errors.New("用户名或者密码错误!")
//	}
//	return nil
//}
