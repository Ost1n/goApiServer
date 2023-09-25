package main

import (
	"training-go-ostin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	messageRouter := router.MessageRouter{}
	replyRouter := router.ReplyRouter{}

	messageRouter.Init(r)
	replyRouter.Init(r)

	r.Run("127.0.0.1:3000")
}
