package controller

import (
	"ace/model"
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfoController(c *gin.Context) {
	var request service.GetUserInfo
	if err := c.Bind(&request); err == nil {
		user, ok := c.Get("user")
		if !ok {
			serializer.NeedLogin(c)
			c.Abort()
			return
		}
		res := request.GetUserInfo(user.(model.User))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}
