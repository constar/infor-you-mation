package main

import (
	"flag"
	"github.com/aszxqw/igo"
	"github.com/golang/glog"
	"time"
)

func main() {
	flag.Parse()
	for {
		content := igo.HttpGet("http://bbs.byr.cn/rss/board-Advertising")
		content = convert(content)
		msgs := Parse(content)
		if msgs == nil {
			glog.Error("Parse failed")
		} else {
			for _, item := range msgs {
				glog.Info(item.GetTitle())
				glog.Info(item.GetUrl())
			}
		}
		println("sleep")
		time.Sleep(5 * time.Second)
	}
}
