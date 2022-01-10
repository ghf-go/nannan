package commonlogic

import (
	"github.com/ghf-go/nannan/mod"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
)

//保存数据到es
func EsSaveFeed(feed *EsFeed) bool {
	r, e := logic.GetEsClient().Update(_esFeed, strconv.FormatInt(feed.FeedId, 10), feed)
	if e != nil {
		mod.Error("ES 更新 %s %s 失败 %s ", _esFeed, feed.FeedId, e.Error())
		return false
	}
	return r.IsSuccess()
}

//从es中查找用户信息
func EsFindUserInfo(feedid int64) *EsFeed {
	r := &esFindFeed{}
	e := logic.GetEsClient().Find(_esFeed, strconv.FormatInt(feedid, 10), r)
	if e != nil {
		mod.Debug("ES 查询 %s 失败 %s", _esFeed, e.Error())
		return nil
	}
	if r.Found {
		return r.Data
	}
	return nil
}

// Es 删除账号信息
func EsDelUserInfo(feedid int64) bool {
	r, e := logic.GetEsClient().Delete(_esFeed, strconv.FormatInt(feedid, 10))
	if e != nil {
		mod.Error("ES 删除 %s  %d 失败 %s ", _esFeed, feedid, e.Error())
		return false
	}

	return r.IsSuccess()
}
