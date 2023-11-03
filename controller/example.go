package controller

import (
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Example(c *gin.Context) {
	var request service.Example
	if err := c.Bind(&request); err == nil {
		res := request.Example()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}
