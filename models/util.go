package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

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

func GetPastDayRange() bson.M {
	righttime := time.Now()
	d, err := time.ParseDuration("-24h")
	if err != nil {
		panic(err)
	}
	lefttime := righttime.Add(d)
	return bson.M{"$gte": lefttime, "$lt": righttime}
}
