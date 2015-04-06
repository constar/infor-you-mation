package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

func GetTopicTrend(keyword string) []int {
	n := 7
	days := GetPastDayRanges(n)
	cnts := make([]int, n)
	for i := 0; i < len(days); i++ {
		cnts[i] = GetTopicDayTrend(keyword, days[i])
	}
	return cnts
}

func GetTopicDayTrend(keyword string, dayrange bson.M) int {
	c := client.dbSess.DB(client.dbName).C(keywordColName)
	m1 := bson.M{"$match": bson.M{"lastmodified": dayrange, "keyword": keyword}}
	m2 := bson.M{"$group": bson.M{"_id": "$keyword", "count": bson.M{"$sum": 1}}}
	pipe := c.Pipe([]bson.M{m1, m2})
	var results []struct {
		Id    string "_id"
		Count int
	}
	err := pipe.All(&results)
	if err != nil {
		beego.Error(err)
		return 0
	}
	if len(results) != 1 {
		beego.Error(results)
		return 0
	}
	return results[0].Count
}
