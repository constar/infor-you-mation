package main

import (
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

const (
	KeywordCol = "keyword"
)

type KeywordDispatcher struct {
	trie   *igo.Trie
	lock   sync.RWMutex
	dbSess *mgo.Session
}

type KeywordColItem struct {
	Id      bson.ObjectId "_id"
	Keyword string
	Feedid  bson.ObjectId
}

func NewKeywordDispatcher() *KeywordDispatcher {
	kw := new(KeywordDispatcher)
	kw.trie = igo.NewTrie()
	var err error
	kw.dbSess, err = mgo.Dial(MongoDBHost)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return kw
}

func (kw *KeywordDispatcher) Insert(word string) {
	kw.lock.Lock()
	defer kw.lock.Unlock()
	if err := kw.trie.Insert(word); err != nil {
		glog.Error(err)
	}
}

func (kw *KeywordDispatcher) Dispatch(text string, feedid bson.ObjectId) {
	kw.lock.RLock()
	defer kw.lock.RUnlock()
	res := kw.trie.Find(text)
	for i := 0; i < len(res); i++ {
		glog.Info(res[i].Pattern)
		err := kw.dispatchOne(res[i].Pattern, feedid)
		if err != nil {
			glog.Error(err)
		}
	}
}

func (kw *KeywordDispatcher) dispatchOne(keyword string, feedid bson.ObjectId) error {
	c := kw.dbSess.DB(MongoDBHost).C(KeywordCol)
	kci := KeywordColItem{
		bson.NewObjectId(),
		keyword,
		feedid,
	}
	return c.Insert(&kci)
}
