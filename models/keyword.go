package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	keywordColName = "keyword"
)

type KeywordFeedPair struct {
	Id      bson.ObjectId "_id"
	Keyword string        "keyword"
	Feedid  bson.ObjectId "feedid"
}

func GetKeywordFeedPairs(word string, n int) ([]KeywordFeedPair, error) {
	c := client.dbSess.DB(client.dbName).C(keywordColName)
	query := bson.M{"keyword": word}
	kwfp := make([]KeywordFeedPair, 0)
	err := c.Find(query).Limit(n).Sort("-_id").All(&kwfp)
	return kwfp, err
}
