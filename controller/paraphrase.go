package controller

import (
	"ace/model"
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParaphraseGenerator(c *gin.Context) {
	var request service.ParaphraseGeneratorRequest
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

func ParaphraseHistory(c *gin.Context) {
	var request service.ParaphraseHistoryRequest
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

func ParaphraseHistoryById(c *gin.Context) {
	var request service.ParaphraseHistoryIdRequest
	if err := c.Bind(&request); err == nil {
		user, ok := c.Get("user")
		if !ok {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		res := request.GetToneContentById(user.(model.User))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}
