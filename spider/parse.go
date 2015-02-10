package main

import (
	"encoding/xml"
	"github.com/golang/glog"
	//"io/ioutil"
)

type ByrRss struct {
	//Version string    `xml:"version,attr"`
	ByrRssChannel ByrRssChannel `xml:"channel"`
}

type ByrRssChannel struct {
	Title         string `xml:"title"`
	Description   string `xml:"description"`
	Link          string `xml:"link"`
	Language      string `xml:"language"`
	Generator     string `xml:"generator"`
	LastBuildDate string `xml:"lastBuildDate"`
	Items         []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Comments    string `xml:"comments"`
	Description string `xml:"description"`
}

type ByrMessage struct {
	title   string
	content string
	url     string
}

func (bm *ByrMessage) GetTitle() string {
	return bm.title
}

func (bm *ByrMessage) GetContent() string {
	return bm.content
}

func (bm *ByrMessage) GetUrl() string {
	return bm.url
}

func Parse(content []byte) []Message {
	if content == nil {
		glog.Error("content is nil")
		return nil
	}
	var rss ByrRss
	err := xml.Unmarshal(content, &rss)
	if err != nil {
		glog.Error(err)
		return nil
	}
	//fmt.Println(rss)
	size := len(rss.ByrRssChannel.Items)
	msgs := make([]Message, 0, size)
	for i := 0; i < size; i++ {
		msg := ByrMessage{
			title:   rss.ByrRssChannel.Items[i].Title,
			content: rss.ByrRssChannel.Items[i].Description,
			url:     rss.ByrRssChannel.Items[i].Link,
		}
		msgs = append(msgs, &msg)
		//println(rss.ByrRssChannel.Items[i].Title)
	}
	return msgs
}

//func main() {
//	content, err := ioutil.ReadFile("1.xml")
//	if err != nil {
//		glog.Fatal(err)
//	}
//
//	content = []byte(convert(string(content)))
//
//	var rss ByrRss
//}
