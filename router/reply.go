package router

import (
	"net/http"
	"time"
	"training-go-ostin/entities"

	"github.com/gin-gonic/gin"
)

type ReplyRouter struct{}

type ParamReplyCreate struct {
	Content  string `json:"content" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UriParamReplyCreate struct {
	MessageId int `uri:"messageId"  binding:"required"`
}

func (re *ReplyRouter) Init(r *gin.Engine) {
	r.POST("/api/message/:messageId/reply", re.Create)
}

func (m *ReplyRouter) Create(c *gin.Context) {
	var param ParamReplyCreate
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var uriParam UriParamReplyCreate
	if err := c.BindUri(&uriParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := &entities.MessageEntity{}

	if err := message.GetById(uriParam.MessageId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message not found"})
		return
	}

	reply := entities.ReplyEntity{
		MessageId: uriParam.MessageId,
		Content:   param.Content,
		Username:  param.Username,
		CreatedAt: time.Now(),
	}
	reply.Save()

	c.JSON(200, reply)
}
