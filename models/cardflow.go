package models

import (
	"github.com/golang/glog"
	"strings"
)

type CardFlow struct {
	Cards []*Card
}

type Card struct {
	Topic string
	Feeds []*Feed
}

const FeedsLimit = 5
const CardFlowN = 2

var Topics = [...]string{
	"实习",
	"兼职",
	"设计",
	"前端",
	"PHP",
	"C++",
	"iOS",
	"Android",
}

func GetCardFlows() []*CardFlow {
	cf := make([]*CardFlow, CardFlowN)
	for i := 0; i < len(cf); i++ {
		cf[i] = &CardFlow{make([]*Card, 0)}
	}
	for i := 0; i < len(Topics); i++ {
		c := GetCardByTopic(Topics[i])
		if c != nil {
			cf[i%len(cf)].Cards = append(cf[i%len(cf)].Cards, c)
		}
	}
	return cf
}

func GetCardByTopic(topic string) *Card {
	kfps, err := GetKeywordFeedPairs(strings.ToLower(topic), FeedsLimit)
	if err != nil {
		glog.Error(err)
		return nil
	}
	c := Card{topic, make([]*Feed, 0)}
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
