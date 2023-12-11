package service

import (
	"ace/pkg"
	"ace/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mime/multipart"
	"time"
)

type Upload struct {
	File *multipart.FileHeader `json:"file"`
}

func (u *Upload) Store(c *gin.Context) serializer.Response {
	filename := fmt.Sprintf("file-%v-%v", time.Now().Second(), u.File.Filename)
	err := c.SaveUploadedFile(u.File, fmt.Sprintf("%v/%v", pkg.Upload.Path, filename))
	if err != nil {
		zap.L().Error("[Upload] Store file failure", zap.Error(err))
		return serializer.StoreFileError(err)
	}
	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
