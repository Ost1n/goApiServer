package entities

import (
	"time"
)

var messageAutoIncreacementId = 1
var messages []MessageEntity

type MessageEntity struct {
	Id        int           `json:"id"`
	Content   string        `json:"content"`
	Username  string        `json:"username"`
	CreatedAt time.Time     `json:"createdAt"`
	Replys    []ReplyEntity `json:"replys"`
}

func (m *MessageEntity) GetMessages() []MessageEntity {
	if messages == nil {
		return []MessageEntity{}
	}

	return messages
}

func (m *MessageEntity) Save() {
	m.Id = messageAutoIncreacementId
	m.Replys = []ReplyEntity{}
	messages = append(messages, *m)
	messageAutoIncreacementId++
}

func (m *MessageEntity) GetById(id int) error {
	for i, ele := range messages {
		if ele.Id == id {
			*m = messages[i]

			return nil
		}
	}

	return &EntityError{}
}

func (m *MessageEntity) Modify(content string) {
	for i, ele := range messages {
		if ele.Id == m.Id {
			messages[i].Content = content
			*m = messages[i]

			return
		}
	}
}

func (m *MessageEntity) Delete() {
	for i, ele := range messages {
		if ele.Id == m.Id {
			messages = append(messages[:i], messages[i+1:]...)

			return
		}
	}
}
