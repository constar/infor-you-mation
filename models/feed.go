package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	feedColName = "feeds"
)

type BriefFeed struct {
	Title string "title"
	Url   string "url"
}

type Feed struct {
	Id      bson.ObjectId "_id"
	Title   string        "title"
	Content string        "content"
	Url     string        "url"
	Urlmd5  string        "urlmd5"
}

func GetBriefFeedById(oid bson.ObjectId) ([]BriefFeed, error) {
	c := client.dbSess.DB(client.dbName).C(feedColName)
	query := bson.M{"_id": oid}
	feeds := make([]BriefFeed, 0)
	err := c.Find(query).All(&feeds)
	return feeds, err
}
