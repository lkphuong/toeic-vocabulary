package utils

import "github.com/gin-gonic/gin"

type Response struct {
	HttpCode  int         `json:"-"`
	ErrorCode int         `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Errors    interface{} `json:"errors"`
}

func JSONResponse(res Response, g *gin.Context) {
	g.JSON(res.HttpCode, gin.H{
		"error_code": res.ErrorCode,
		"message":    res.Message,
		"data":       res.Data,
		"errors":     res.Errors,
	})
}

func NewResponse(
	data interface{},
	message string,
	codes ...int,
) *Response {

	httpCode := 200        // default value
	errorCode := 0         // default value
	var errors interface{} // default value is nil

	if len(codes) > 0 {
		httpCode = codes[0]
	}
	if len(codes) > 1 {
		errorCode = codes[1]
	}
	if len(codes) > 2 {
		errors = codes[2]
	}

	return &Response{
		HttpCode:  httpCode,
		ErrorCode: errorCode,
		Message:   message,
		Data:      data,
		Errors:    errors,
	}
}
