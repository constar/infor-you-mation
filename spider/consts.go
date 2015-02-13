package main

const (
	DBName      = "inforyoumation"
	MongoDBHost = "127.0.0.1"
)

var (
	RssUrls = [...]string{
		"http://bbs.byr.cn/rss/board-WorkLife",
		"http://bbs.byr.cn/rss/board-Advertising",
	}
	Keywords = [...]string{
		"实习",
		"兼职",
		"全职",
	}
)
