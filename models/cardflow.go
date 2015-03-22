package models

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

type CardFlow struct {
	Cards []*Card
}

type Card struct {
	Topic        string
	YesterdayAdd int
	Feeds        []*Feed
}

const CardFlowN = 2

var Topics = [...]string{
	"实习/兼职",
	"大数据",
	"人工智能",
	"设计",
	"前端",
	"PHP",
	"C++",
	"iOS",
	"Android",
}

func GetRandomCardFlows(row_num int) []*CardFlow {
	cf := make([]*CardFlow, CardFlowN)
	for i := 0; i < len(cf); i++ {
		cf[i] = &CardFlow{make([]*Card, 0)}
	}
	rnd := rand.Perm(len(Topics))
	for _, i := range rnd {
		c := GetCardByTopic(Topics[i], row_num)
		if c != nil {
			cf[i%len(cf)].Cards = append(cf[i%len(cf)].Cards, c)
		}
	}
	return cf
}

func GetHotCardFlows(row_num int) []*CardFlow {
	topics := GetHotTopics()
	flows := make([]*CardFlow, CardFlowN)
	for i := 0; i < len(flows); i++ {
		flows[i] = &CardFlow{make([]*Card, 0)}
	}
	for i, topic := range topics {
		j := i % len(flows)
		c := GetCardByTopic(topic, row_num)
		if c != nil {
			flows[j].Cards = append(flows[j].Cards, c)
		}
	}
	return flows
}

func GetHotTopics() (topics []string) {
	c := client.dbSess.DB(client.dbName).C(keywordColName)
	pastday := GetPastDayRange()
	m1 := bson.M{"$match": bson.M{"lastmodified": pastday}}
	m2 := bson.M{"$group": bson.M{"_id": "$keyword", "count": bson.M{"$sum": 1}}}
	m3 := bson.M{"$sort": bson.M{"count": -1}}
	pipe := c.Pipe(
		[]bson.M{
			m1,
			m2,
			m3,
		})
	var results []struct {
		Id    string "_id"
		Count int
	}
	err := pipe.All(&results)
	if err != nil {
		glog.Error(err)
		return
	}
	for _, res := range results {
		topics = append(topics, res.Id)
	}
	return
}

func GetCardFlows(row_num int) []*CardFlow {
	cf := make([]*CardFlow, CardFlowN)
	for i := 0; i < len(cf); i++ {
		cf[i] = &CardFlow{make([]*Card, 0)}
	}
	for i := 0; i < len(Topics); i++ {
		c := GetCardByTopic(Topics[i], row_num)
		j := i % len(cf)
		if c != nil {
			cf[j].Cards = append(cf[j].Cards, c)
		}
	}
	return cf
}

func GetCardByTopic(topic string, row_num int) *Card {
	kfps, err := GetKeywordFeedPairs(topic, row_num)
	if err != nil {
		glog.Error(err)
		return nil
	}
	c := Card{topic, GetYesterdayAddByKeyword(topic), make([]*Feed, 0)}
	for i := 0; i < len(kfps); i++ {
		feeds, err := GetFeedById(kfps[i].Feedid)
		if err != nil {
			glog.Error(err)
			return nil
		}
		if len(feeds) != 1 {
			glog.Error("feeds illegal!!!")
			return nil
		}
		c.Feeds = append(c.Feeds, &feeds[0])
	}
	return &c
}
