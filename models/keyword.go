package models

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	keywordColName = "keyword"
)

type KeywordFeedPair struct {
	Id      bson.ObjectId "_id"
	Keyword string        "keyword"
	Feedid  bson.ObjectId "feedid"
}

func (this *KeywordFeedPair) String() string {
	return fmt.Sprintf("Id:%v Keyword:%v Feedid:%v", this.Id, this.Keyword, this.Feedid)
}

func GetKeywordFeedPairs(word string, n int) ([]KeywordFeedPair, error) {
	c := client.dbSess.DB(client.dbName).C(keywordColName)
	query := bson.M{"keyword": word}
	kwfp := make([]KeywordFeedPair, 0)
	err := c.Find(query).Limit(n).Sort("-_id").All(&kwfp)
	return kwfp, err
}

func GetYesterdayAddByKeyword(word string) int {
	c := client.dbSess.DB(client.dbName).C(keywordColName)
	righttime := time.Now()
	d, err := time.ParseDuration("-24h")
	if err != nil {
		glog.Error(err)
		return 0
	}
	lefttime := righttime.Add(d)

	query := bson.M{"keyword": word, "lastmodified": bson.M{"$gte": lefttime, "$lt": righttime}}
	fmt.Println(lefttime)
	fmt.Println(righttime)
	cnt, err := c.Find(query).Count()
	if err != nil {
		glog.Error(err)
		return 0
	}
	return cnt
}
