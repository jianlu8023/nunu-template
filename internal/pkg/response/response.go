package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: ErrSuccess.Code, Message: ErrSuccess.Message, Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: errorCodeMap[err.Error()], Message: err.Error(), Data: data}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[string]int{}

func newError(code int, msg string) *Error {
	err := errors.New(msg)
	errorCodeMap[err.Error()] = code
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) ErrorCode() int {
	return e.Code
}

func (e Error) SetError(err string) *Error {
	e.Message = err
	return &e
}
