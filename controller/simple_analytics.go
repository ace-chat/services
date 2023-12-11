package controller

import (
	"ace/serializer"
	"ace/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GeneratorSimpleAnalytics(c *gin.Context) {
	var request service.SimpleAnalytics
	if err := c.Bind(&request); err == nil {
		res := request.Generator()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError(err))
	}
}
