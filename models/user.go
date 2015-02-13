package models

import (
	"fmt"
	"github.com/golang/glog"
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

func ValidateUser(username string, passwd string) bool {
	query := bson.M{"username": username, "password": passwd}
	c := connOption.dbSess.DB(connOption.dbName).C(userColName)
	cnt, err := c.Find(query).Count()
	if err != nil {
		glog.Error(err)
		return false
	}
	if cnt == 1 {
		return true
	} else {
		glog.Error("username: ", username, " password ", passwd, " not found.")
		return false
	}
}

//func ValidateUser(user User) error {
//	orm := getOrm()
//	var u User
//	orm.Where("username=? and pwd=?", user.Username, user.Password).Find(&u)
//	if u.Username == "" {
//		return errors.New("用户名或者密码错误!")
//	}
//	return nil
//}
