package main

import (
	"github.com/yanyiwu/igo"
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
	host   string
	dbSess *mgo.Session
	dbName string
}

var (
	connOption ConnectOption
)

func Connect(host string, dbName string) error {
	var err error
	connOption.host = host
	connOption.dbSess, err = mgo.Dial(host)
	connOption.dbName = dbName
	return err
}

func Close() {
	connOption.dbSess.Close()
}

func Insert(
	colName string,
	title string,
	content string,
	url string) error {
	c := connOption.dbSess.DB(connOption.dbName).C(colName)

	bdi := ByrDataItem{
		bson.NewObjectId(),
		title,
		content,
		url,
		igo.GetMd5String(url),
	}

	return c.Insert(&bdi)
}
