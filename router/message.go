package router

import (
	"net/http"
	"time"
	"training-go-ostin/entities"

	"github.com/gin-gonic/gin"
)

type MessageRouter struct{}

type ParamMessageCreate struct {
	Content  string `json:"content" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UriParamMessageModify struct {
	Id int `uri:"id"  binding:"required"`
}

type ParamMessageModify struct {
	Content string `json:"content,required"`
}

type UriParamMessageDelete struct {
	Id int `uri:"id"  binding:"required"`
}

func (m *MessageRouter) Init(r *gin.Engine) {
	r.GET("/api/message/list", m.List)
	r.POST("/api/message", m.Create)
	r.PUT("/api/message/:id", m.Modify)
	r.DELETE("/api/message/:id", m.Delete)
}

func (m *MessageRouter) List(c *gin.Context) {
	messageEntity := entities.MessageEntity{}
	replyEntity := entities.ReplyEntity{}
	messages := messageEntity.GetMessages()

	for i, ele := range messages {
		messages[i].Replys = replyEntity.GetReplys(ele.Id)
	}

	c.JSON(200, messages)
}

func (m *MessageRouter) Create(c *gin.Context) {
	var param ParamMessageCreate
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := entities.MessageEntity{
		Content:   param.Content,
		Username:  param.Username,
		CreatedAt: time.Now(),
	}
	message.Save()

	c.JSON(200, message)
}

func (m *MessageRouter) Modify(c *gin.Context) {
	var param ParamMessageModify
	var uriParam UriParamMessageModify

	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindUri(&uriParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := &entities.MessageEntity{}

	if err := message.GetById(uriParam.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message not found"})
		return
	}

	message.Modify(param.Content)

	c.JSON(200, message)
}

func (m *MessageRouter) Delete(c *gin.Context) {
	var uriParam UriParamMessageModify

	if err := c.BindUri(&uriParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := &entities.MessageEntity{}

	if err := message.GetById(uriParam.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message not found"})
		return
	}

	message.Delete()

	c.JSON(200, message)
}
