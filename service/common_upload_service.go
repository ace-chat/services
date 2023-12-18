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
	File *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}

func (u *Upload) Store(c *gin.Context) serializer.Response {
	filename := fmt.Sprintf("file-%v-%v", time.Now().Unix(), u.File.Filename)
	err := c.SaveUploadedFile(u.File, fmt.Sprintf("%v/%v", pkg.Upload.Path, filename))
	if err != nil {
		zap.L().Error("[Upload] Store file failed", zap.Error(err))
		return serializer.StoreFileError(err)
	}
	return serializer.Response{
		Code: 200,
		Data: filename,
	}
}
