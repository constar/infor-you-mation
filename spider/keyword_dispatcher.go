package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"sync"
)

var dispatcher *KeywordDispatcher

func init() {
	dispatcher = NewKeywordDispatcher()
	if dispatcher == nil {
		panic("NewKeywordDispatcher failed")
	}
	for i := 0; i < len(Keywords); i++ {
		k := strings.ToLower(Keywords[i])
		dispatcher.Insert(k)
	}
}

func Dispatch(text string, feedid bson.ObjectId) {
	dispatcher.Dispatch(text, feedid)
}

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

func (kci *KeywordColItem) String() string {
	return fmt.Sprintf("%v %s %v", kci.Id, kci.Keyword, kci.Feedid)
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
	text = strings.ToLower(text)
	res := kw.trie.Find(text)
	for i := 0; i < len(res); i++ {
		err := kw.dispatchOne(res[i].Pattern, feedid)
		if err != nil {
			glog.Info(err)
		}
	}
}

func (kw *KeywordDispatcher) dispatchOne(keyword string, feedid bson.ObjectId) error {
	c := kw.dbSess.DB(DBName).C(KeywordCol)
	kci := KeywordColItem{
		bson.NewObjectId(),
		keyword,
		feedid,
	}
	glog.Infof("insert %s to %s.%s", kci, DBName, KeywordCol)
	return c.Insert(&kci)
}
