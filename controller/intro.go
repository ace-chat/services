package controller

import (
	"ace/model"
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IntroGenerator(c *gin.Context) {
	var request service.IntroGeneratorRequest
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

func IntroHistory(c *gin.Context) {
	var request service.IntroHistoryRequest
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

func IntroHistoryById(c *gin.Context) {
	var request service.IntroHistoryIdRequest
	if err := c.Bind(&request); err == nil {
		user, ok := c.Get("user")
		if !ok {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		res := request.GetIntroContentById(user.(model.User))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}
