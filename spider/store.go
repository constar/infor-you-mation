package main

import (
	"github.com/aszxqw/igo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ByrDataItem struct {
	Id      bson.ObjectId "_id"
	Title   string
	Content string
	Url     string
	UrlMd5  string
}

type ConnectOption struct {
	host    string
	dbSess  *mgo.Session
	dbName  string
	colName string
}

var (
	connOption ConnectOption
)

func Connect(host string, dbName string, colName string) error {
	var err error
	connOption.host = host
	connOption.dbSess, err = mgo.Dial(host)
	connOption.dbName = dbName
	connOption.colName = colName
	return err
}

func Close() {
	connOption.dbSess.Close()
}

func Insert(
	title string,
	content string,
	url string) error {
	c := connOption.dbSess.DB(connOption.dbName).C(connOption.colName)

	bdi := ByrDataItem{
		bson.NewObjectId(),
		title,
		content,
		url,
		igo.GetMd5String(url),
	}

	return c.Insert(&bdi)
}
