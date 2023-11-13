package controller

import (
	"ace/model"
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommonPlatforms(c *gin.Context) {
	var request service.PlatformRequest
	if err := c.Bind(&request); err == nil {
		res := request.GetPlatforms()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func CommonRegion(c *gin.Context) {
	var request service.RegionRequest
	if err := c.Bind(&request); err == nil {
		res := request.GetRegions()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func CommonTones(c *gin.Context) {
	var request service.TonesRequest
	if err := c.Bind(&request); err == nil {
		res := request.GetTones()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func CommonVoices(c *gin.Context) {
	var request service.CommonVoicesRequest
	if err := c.Bind(&request); err == nil {
		user, ok := c.Get("user")
		if !ok {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		res := request.GetVoices(user.(model.User))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func CommonCreateVoice(c *gin.Context) {
	var request service.CommonCreateVoiceRequest
	if err := c.Bind(&request); err == nil {
		user, ok := c.Get("user")
		if !ok {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		res := request.CreateVoice(user.(model.User))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func CommonLanguages(c *gin.Context) {
	var request service.LanguageRequest
	if err := c.Bind(&request); err == nil {
		res := request.GetLanguages()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}

func CommonGender(c *gin.Context) {
	var request service.GenderRequest
	if err := c.Bind(&request); err == nil {
		res := request.GetGender()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}
