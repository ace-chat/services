package controller

import (
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request service.Login
	if err := c.Bind(&request); err == nil {
		res := request.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.ParamError(err))
	}
}
