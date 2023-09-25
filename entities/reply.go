package entities

import (
	"time"
)

var autoIncreacementId = 1
var replys []ReplyEntity

type ReplyEntity struct {
	Id        int `json:"id"`
	MessageId int
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

func (r *ReplyEntity) GetReplys(messageId int) []ReplyEntity {
	ret := []ReplyEntity{}
	for _, ele := range replys {
		if ele.MessageId == messageId {
			ret = append(ret, ele)
		}
	}
	return ret
}

func (r *ReplyEntity) Save() {
	r.Id = autoIncreacementId
	replys = append(replys, *r)
	autoIncreacementId++
}
