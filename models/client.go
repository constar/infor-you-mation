package models

import (
	"gopkg.in/mgo.v2"
)

const (
	mongoHost   = "127.0.0.1"
	mongoDBName = "inforyoumation"
)

var client MongoClient

type MongoClient struct {
	host   string
	dbSess *mgo.Session
	dbName string
}

func Connect(host string, dbName string) error {
	var err error
	client.host = host
	client.dbSess, err = mgo.Dial(host)
	client.dbName = dbName
	return err
}

func init() {
	err := Connect(mongoHost, mongoDBName)
	if err != nil {
		panic(err)
	}
}
