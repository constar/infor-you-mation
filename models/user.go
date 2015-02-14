package models

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

const (
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

func RegisterUser(username string, passwd string) error {
	c := client.dbSess.DB(client.dbName).C(userColName)
	u := User{
		bson.NewObjectId(),
		username,
		passwd,
	}
	return c.Insert(&u)
}

func ValidateUser(username string, passwd string) bool {
	query := bson.M{"username": username, "password": passwd}
	c := client.dbSess.DB(client.dbName).C(userColName)
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
