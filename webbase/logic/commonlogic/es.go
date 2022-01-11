package commonlogic

import (
	"github.com/ghf-go/nannan/mod"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
)

//保存数据到es
func EsSaveFeed(feed *EsFeed) bool {
	return logic.GetEsClient().Update(_esFeed, strconv.FormatInt(feed.FeedId, 10), feed)
}

//从es中查找用户信息
func EsFindUserInfo(feedid int64) *EsFeed {
	r := &EsFeed{}
	e := logic.GetEsClient().Find(_esFeed, strconv.FormatInt(feedid, 10), r)
	if e != nil {
		mod.Debug("ES 查询 %s 失败 %s", _esFeed, e.Error())
		return nil
	}

	return r
}

// Es 删除账号信息
func EsDelUserInfo(feedid int64) bool {
	return logic.GetEsClient().Delete(_esFeed, strconv.FormatInt(feedid, 10))
}
