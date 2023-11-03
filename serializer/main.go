package serializer

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

const (
	CodeNeedLogin     = 20005
	CodeDBError       = 20006
	CodeParamError    = 20007
	CodeNotFoundError = 20008
	CodePasswordError = 20009
	CodeTokenError    = 20009
	CodeExistsError   = 20010
	CodeBotError      = 20011
)

func NeedLogin(c *gin.Context) {
	c.JSON(401, Response{
		Code:    CodeNeedLogin,
		Message: "The user is not logged in, please log in and then perform another operation",
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

func ExistsError() Response {
	msg := "User google key already exists"
	return Err(CodeExistsError, msg, errors.New("user google key already exists"))
}

func RunError(err error) Response {
	msg := "Run Error"
	return Err(CodeBotError, msg, err)
}
