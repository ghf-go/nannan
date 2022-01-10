package commentlogic

import (
	"encoding/json"
	"github.com/ghf-go/nannan/webbase/logic/accountlogic"
	"time"
)

type CommentModel struct {
	ID         int64     `column:"id" json:"id"`
	UserId     int64     `column:"user_id" json:"user_id"`
	TargetId   int64     `column:"target_id" json:"target_id"`
	ParentId   int64     `column:"parent_id" json:"parent_id"`
	TargetType int       `column:"target_type" json:"target_type"`
	Content    string    `column:"content" json:"content"`
	ReplyCount int       `column:"reply_count" json:"reply_count"`
	CreateAt   time.Time `column:"create_at" json:"create_at"`
}

type FeedDetailModel struct {
	ID          int64                  `column:"id" json:"feed_id"`
	UserId      int64                  `column:"user_id" json:"user_id"`
	FeedType    int                    `column:"feed_type" json:"feed_type"`
	FeedImgBase string                 `column:"feed_imgs" json:"-"`
	FeedImgs    []string               `json:"feed_imgs"`
	X           float64                `column:"x" json:"x"`
	Y           float64                `column:"y" json:"y"`
	City        string                 `column:"city" json:"city"`
	ExtBase     string                 `column:"ext" json:"-"`
	ContentBase string                 `column:"content" json:"-"`
	Ext         map[string]interface{} `json:"ext"`
	Content     map[string]interface{} `json:"content"`
	UserInfo    interface{}            `json:"user_info"`
	CreateAt    time.Time              `column:"create_at" json:"create_at"`
}

func (obj *FeedDetailModel) build() {
	json.Unmarshal([]byte(obj.ExtBase), &obj.Ext)
	json.Unmarshal([]byte(obj.ContentBase), &obj.Content)
	json.Unmarshal([]byte(obj.FeedImgBase), &obj.FeedImgs)
	obj.UserInfo = accountlogic.GetProfileByUid(obj.UserId)
}

func (obj *FeedDetailModel) addUser() {
	obj.UserInfo = accountlogic.GetProfileByUid(obj.UserId)
}
