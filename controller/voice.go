package controller

import (
	"ace/model"
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VoiceGenerator(c *gin.Context) {
	var request service.VoiceGeneratorRequest
	if err := c.Bind(&request); err == nil {
		user, ok := c.Get("user")
		if !ok {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		res := request.Generator(user.(model.User))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func VoiceHistory(c *gin.Context) {
	var request service.VoiceHistoryRequest
	if err := c.Bind(&request); err == nil {
		user, ok := c.Get("user")
		if !ok {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		res := request.GetHistory(user.(model.User))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func VoiceHistoryById(c *gin.Context) {
	// TODO voice history api
	//var request service.ParaphraseHistoryIdRequest
	//if err := c.Bind(&request); err == nil {
	//	user, ok := c.Get("user")
	//	if !ok {
	//		serializer.NeedLogin(c)
	//		c.Abort()
	//		return
	//	}
	//	res := request.GetToneContentById(user.(model.User))
	//	c.JSON(http.StatusOK, res)
	//} else {
	//	c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	//}
}
