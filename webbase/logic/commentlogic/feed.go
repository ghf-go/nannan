package commentlogic

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
)

func NewFeed(userid int64, FeedType int, FeedDesc, FeedImgs string, x, y float64, city, ext, content string) {
	id := logic.GetTable(tb_comment_feed).InsertMap(def.Data{
		"user_id":   userid,
		"feed_type": FeedType,
		"feed_desc": FeedDesc,
		"feed_imgs": FeedImgs,
		"x":         x,
		"y":         y,
		"city":      city,
		"ext":       ext,
		"content":   content,
	})
	if id > 0 && logic.IsEsEnable() {
		updateEs(id)
	}

}

func MyPublishFeedList(uid int64, start, limit int) {
	if logic.IsEsEnable() {
		logic.GetEsClient().NewQuery(es_comment_feed).Size(limit).Start(start).MustMatch("user_id", uid)
	}
}

//更新内容到es中
func updateEs(id int64) {
	go func() {
		obj := &FeedDetailModel{}
		if logic.CreateQuery(tb_comment_feed).Where("id=?", id).Frist(obj) == nil {
			obj.build()
			logic.GetEsClient().Update(es_comment_feed, strconv.FormatInt(id, 10), obj)
		}
	}()
}
