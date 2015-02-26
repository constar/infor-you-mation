package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/yanyiwu/infor-you-mation/models"
	"strings"
)

type IndexController struct {
	beego.Controller
}

func (mc *IndexController) Get() {
	mc.Data["CardFlows"] = mc.getCardFlows()
	mc.TplNames = "index.tpl"
}

type CardFlow struct {
	Cards []*Card
}

type Card struct {
	Topic string
	Feeds []*Feed
}

type Feed struct {
	Title string
	Url   string
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
}

func (mc *IndexController) getCardFlows() []*CardFlow {
	cf := make([]*CardFlow, CardFlowN)
	for i := 0; i < len(cf); i++ {
		cf[i] = &CardFlow{make([]*Card, 0)}
	}
	for i := 0; i < len(Topics); i++ {
		c := mc.getCardByTopic(strings.ToLower(Topics[i]))
		if c != nil {
			cf[i%len(cf)].Cards = append(cf[i%len(cf)].Cards, c)
		}
	}
	return cf
}

func (mc *IndexController) getCardByTopic(topic string) *Card {
	kfps, err := models.GetKeywordFeedPairs(topic, FeedsLimit)
	if err != nil {
		glog.Error(err)
		return nil
	}
	c := Card{topic, make([]*Feed, 0)}
	for i := 0; i < len(kfps); i++ {
		feeds, err := models.GetFeedById(kfps[i].Feedid)
		if err != nil {
			glog.Error(err)
			return nil
		}
		if len(feeds) != 1 {
			glog.Error("feeds illegal!!!")
			return nil
		}
		feed := Feed{feeds[0].Title, feeds[0].Url}
		c.Feeds = append(c.Feeds, &feed)
	}
	return &c
}
