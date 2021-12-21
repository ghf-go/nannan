package commentlogic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/web/webbase/logic"
	"github.com/go-redis/redis/v8"
	"time"
)

func NewFeed(userid int64, FeedTitle string, FeedType int, FeedDesc, FeedImgs string, x, y float64, city, ext, content string) {
	id := logic.GetTable(tb_comment_feed).InsertMap(db.Data{
		"user_id":    userid,
		"feed_title": FeedTitle,
		"feed_type":  FeedType,
		"feed_desc":  FeedDesc,
		"feed_imgs":  FeedImgs,
		"x":          x,
		"y":          y,
		"city":       city,
		"ext":        ext,
	})
	if id > 0 {
		logic.GetTable(tb_comment_feed_content).InsertMap(db.Data{
			"id":      id,
			"content": content,
		})
		loadFeedDetailToRedis(id)
	}

}
func FeedList() {

}
func FollowFeedList(uid int64) {

}
func FeedDetail(targetId int64) *FeedDetailModel {
	d := logic.GetRedis().Get(context.Background(), redisFeedKey(targetId)).String()

	detail := &FeedDetailModel{}
	if json.Unmarshal([]byte(d), detail) != nil {
		return nil
	}
	return detail
}
func loadFeedDetailToRedis(targetId int64) *FeedDetailModel {
	detail := &FeedDetailModel{}
	if logic.CreateQuery(tb_comment_feed).Where("id=?", targetId).Frist(detail) != nil {
		return nil
	}
	r, e := logic.GetTable(tb_comment_feed_content).Query(fmt.Sprintf("SELECT content FROM %s WHERE id=?", tb_comment_feed_content), targetId)
	if e != nil {
		return nil
	}
	defer r.Close()
	if r.Next() {
		content := ""
		if r.Scan(&content) != nil {
			return nil
		}
		detail.build(content)
	}
	sobj, e := json.Marshal(detail)
	if e != nil {
		logic.GetRedis().Set(context.Background(), redisFeedKey(targetId), string(sobj), time.Second*86400)
		logic.GetRedis().ZAdd(context.Background(), redisUserFeedKey(detail.UserId), &redis.Z{
			Score:  float64(detail.CreateAt.UnixNano()),
			Member: detail.ID,
		})
	}

	return detail

}
