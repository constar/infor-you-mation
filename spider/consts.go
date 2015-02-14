package main

const (
	MongoDBHost = "127.0.0.1"
	DBName      = "inforyoumation"
	KeywordCol  = "keyword"
)

var (
	RssUrls = [...]string{
		"http://bbs.byr.cn/rss/board-WorkLife",
		"http://bbs.byr.cn/rss/board-Advertising",
		"http://bbs.byr.cn/rss/board-ParttimeJob",
		"http://bbs.byr.cn/rss/board-JobInfo",
		"http://bbs.byr.cn/rss/board-BookTrade",
	}
	Keywords = [...]string{
		"实习",
		"兼职",
		"全职",
	}
)
