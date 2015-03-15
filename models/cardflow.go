package models

import (
	"github.com/golang/glog"
	"math/rand"
	"strings"
)

type CardFlow struct {
	Cards []*Card
}

type Card struct {
	Topic string
	Feeds []*Feed
}

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

func GetCardFlows(row_num int) []*CardFlow {
	cf := make([]*CardFlow, CardFlowN)
	for i := 0; i < len(cf); i++ {
		cf[i] = &CardFlow{make([]*Card, 0)}
	}
	for i := 0; i < len(Topics); i++ {
		c := GetCardByTopic(Topics[i], row_num)
		if c != nil {
			cf[i%len(cf)].Cards = append(cf[i%len(cf)].Cards, c)
		}
	}
	return cf
}

func GetCardByTopic(topic string, row_num int) *Card {
	kfps, err := GetKeywordFeedPairs(strings.ToLower(topic), row_num)
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
