package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"github.com/yanyiwu/infor-you-mation/models"
)

type LoginController struct {
	beego.Controller
}

func (mc *LoginController) Get() {
	mc.Data["Website"] = "InfoYouMation"
	mc.Data["Email"] = "i@yanyiwu.com"
	mc.TplNames = "login.tpl"
}

type Card struct {
	Topic string
	Feeds []*Feed
}

type Feed struct {
	Title string
}

const FeedsLimit = 5

var Topics = [...]string{
	"实习",
	"兼职",
}

func (mc *LoginController) getCards() []*Card {
	cards := make([]*Card, 0)
	for i := 0; i < len(Topics); i++ {
		c := mc.getCardByTopic(Topics[i])
		if c != nil {
			cards = append(cards, c)
		}
	}
	return cards
}

func (mc *LoginController) getCardByTopic(topic string) *Card {
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
		feed := Feed{feeds[0].Title}
		c.Feeds = append(c.Feeds, &feed)
	}
	return &c
}

func (mc *LoginController) Post() {
	inputs := mc.Input()
	username := inputs.Get("username")
	passwd := inputs.Get("password")
	passwd = igo.GetMd5String(passwd)
	if models.ValidateUser(username, passwd) {
		glog.Info("username:", username, " login success!")
		mc.Data["Cards"] = mc.getCards()
		mc.TplNames = "index.tpl"
	} else {
		mc.TplNames = "loginfailure.tpl"
	}
}
