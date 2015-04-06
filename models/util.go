package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func GetPastDayRange() bson.M {
	return GetPastDayRanges(1)[0]
}

func GetPastDayRanges(n int) []bson.M {
	d, err := time.ParseDuration("-24h")
	if err != nil {
		panic(err)
	}
	ms := make([]bson.M, n)
	righttime := time.Now()
	lefttime := righttime.Add(d)
	for i := 0; i < n; i++ {
		ms[n-i-1] = bson.M{"$gte": lefttime, "$lt": righttime}
		righttime = righttime.Add(d)
		lefttime = righttime.Add(d)
	}
	return ms
}

/* stale function */
func GetYesterdayRange() bson.M {
	righttime := time.Now()
	righttime = time.Date(righttime.Year(), righttime.Month(), righttime.Day(), 0, 0, 0, 0, time.UTC)
	d, err := time.ParseDuration("-24h")
	if err != nil {
		panic(err)
	}
	lefttime := righttime.Add(d)
	return bson.M{"$gte": lefttime, "$lt": righttime}
}
