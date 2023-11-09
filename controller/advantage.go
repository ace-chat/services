package controller

import (
	"ace/model"
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdvantageGenerator(c *gin.Context) {
	var request service.AdvantageGeneratorRequest
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

func AdvantageHistory(c *gin.Context) {
	var request service.AdvantageHistoryRequest
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

func AdvantageHistoryById(c *gin.Context) {
	var request service.AdvantageHistoryIdRequest
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
