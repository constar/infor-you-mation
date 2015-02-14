package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	feedColName = "feeds"
)

type Feed struct {
	Id      bson.ObjectId "_id"
	Title   string        "title"
	Content string        "content"
	Url     string        "url"
	Urlmd5  string        "urlmd5"
}

func GetFeedById(oid bson.ObjectId) ([]Feed, error) {
	c := client.dbSess.DB(client.dbName).C(feedColName)
	query := bson.M{"_id": oid}
	feeds := make([]Feed, 0)
	err := c.Find(query).All(&feeds)
	return feeds, err
}
