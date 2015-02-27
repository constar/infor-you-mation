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
		"http://www.newsmth.net/nForum/rss/board-Career_Campus",
		"http://www.newsmth.net/nForum/rss/board-Career_Investment",
		"http://www.newsmth.net/nForum/rss/board-Career_PHD",
		"http://www.newsmth.net/nForum/rss/board-Career_Plaza",
		"http://www.newsmth.net/nForum/rss/board-Career_Upgrade",
		"http://www.newsmth.net/nForum/rss/board-ExecutiveSearch",
	}
	Keywords = [...]string{
		"实习",
		"兼职",
		"全职",
		"设计",
		"前端",
		"PHP",
		"机器学习",
		"Android",
		"iOS",
		"C++",
		"Java",
		"创业",
		"产品",
	}
)
