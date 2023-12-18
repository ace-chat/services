package serializer

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

const (
	CodeNeedLogin                  = 20005
	CodeDBError                    = 20006
	CodeParamError                 = 20007
	CodeNotFoundError              = 20008
	CodePasswordError              = 20009
	CodeTokenError                 = 20010
	CodeGeneratorError             = 20011
	CodeNotFoundServiceError       = 20012
	CodeNotFoundGeneratorItemError = 20013
	CodeStoreFileError             = 20014
)

func NeedLogin(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    CodeNeedLogin,
		Message: "User login has expired, please log in again",
	})
}

func Err(code int, msg string, err error) Response {
	res := Response{
		Code:    code,
		Message: msg,
	}

	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBError Database error
func DBError(err error) Response {
	msg := "Incorrect amount entered, please try again"
	return Err(CodeDBError, msg, err)
}

func UserNotFoundError(err error) Response {
	msg := "The user was not found"
	return Err(CodeNotFoundError, msg, err)
}

// ParamError Params Error
func ParamError(err error) Response {
	msg := "The parameter is wrong, please confirm the operation process before performing this operation"
	return Err(CodeParamError, msg, err)
}

// PasswordError password error
func PasswordError() Response {
	msg := "Password invalid"
	return Err(CodePasswordError, msg, errors.New("password invalid"))
}

// TokenError generator token error
func TokenError(err error) Response {
	msg := "Error generating login information"
	return Err(CodeTokenError, msg, err)
}

func NotFoundError(err error) Response {
	msg := "User google key already exists"
	return Err(CodeNotFoundError, msg, err)
}

func GeneratorError(err error) Response {
	msg := "Generator ADS failed"
	return Err(CodeGeneratorError, msg, err)
}

func NotFoundPlatformError(err error) Response {
	msg := "platform invalid, please make sure you select the correct platform"
	return Err(CodeNotFoundGeneratorItemError, msg, err)
}

func NotFoundToneError(err error) Response {
	msg := "tone invalid, please make sure you select the correct tone"
	return Err(CodeNotFoundGeneratorItemError, msg, err)
}

func NotFoundVoiceError(err error) Response {
	msg := "brand voice invalid, please make sure you select the correct brand voice"
	return Err(CodeNotFoundGeneratorItemError, msg, err)
}

func NotFoundRegionError(err error) Response {
	msg := "region invalid, please make sure you select the correct region"
	return Err(CodeNotFoundGeneratorItemError, msg, err)
}

func NotFoundLanguageError(err error) Response {
	msg := "language invalid, please make sure you select the correct language"
	return Err(CodeNotFoundGeneratorItemError, msg, err)
}

func NotFoundGenderError(err error) Response {
	msg := "gender invalid, please make sure you select the correct gender"
	return Err(CodeNotFoundGeneratorItemError, msg, err)
}

func NotFoundTypeError(err error) Response {
	msg := "type invalid, please make sure you select the correct type"
	return Err(CodeNotFoundGeneratorItemError, msg, err)
}

func NotFoundServiceError(err error) Response {
	msg := "service invalid, please make sure you select the correct service"
	return Err(CodeNotFoundServiceError, msg, err)
}

func StoreFileError(err error) Response {
	msg := "Failed to operate file"
	return Err(CodeStoreFileError, msg, err)
}
