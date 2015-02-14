package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"sync"
	"time"
)

const (
	SleepSeconds = 60
)

var wg sync.WaitGroup

func spiderRunner(url string) {
	defer wg.Done()
	for {
		content := igo.HttpGet(url)
		content = convert(content)
		msgs := Parse(content)
		if msgs == nil {
			glog.Error("Parse failed")
		} else {
			for _, item := range msgs {
				oid, err := Insert("feeds", item.GetTitle(), item.GetContent(), item.GetUrl())
				if err == nil {
					glog.Info(item.GetTitle(), " ", item.GetUrl())
					Dispatch(item.GetTitle(), oid)
				} else {
					glog.V(2).Info(err)
				}
			}
		}
		glog.V(3).Info("time.Sleep ", SleepSeconds, " seconds")
		time.Sleep(SleepSeconds * time.Second)
	}
}

func main() {
	flag.Parse()
	for i := 0; i < len(RssUrls); i++ {
		url := RssUrls[i]
		wg.Add(1)
		go spiderRunner(url)
		glog.Info(url)
	}
	wg.Wait()
}
