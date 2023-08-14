package format

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// response
// @Description:
//
// @author ght
// @date 2023-08-14 13:49:47
type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Obj     interface{} `json:"obj"`
}

// Success
// @Description:
//
// @author ght
// @date 2023-08-14 13:41:21
//
// @param ctx
// @param message
// @param code
// @param obj
func Success(ctx *gin.Context, message string, code int, obj interface{}) {
	if obj == nil {
		obj = map[string]string{}
	}

	resp := response{
		Success: true,
		Message: message,
		Code:    code,
		Obj:     obj,
	}
	ctx.JSON(http.StatusOK, resp)
}

// Error
// @Description:
//
// @author ght
// @date 2023-08-14 13:41:54
//
// @param ctx
// @param message
// @param code
// @param obj
func Error(ctx *gin.Context, message string, code int, obj interface{}) {
	if obj == nil {
		obj = map[string]string{}
	}
	resp := response{
		Success: false,
		Message: message,
		Code:    code,
		Obj:     obj,
	}
	ctx.JSON(http.StatusOK, resp)
}
