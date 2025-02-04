package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	res := Response{
		Code:    code,
		Data:    err.Error(),
		Message: "error",
	}
	c.JSON(code, res)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	res := Response{
		Code:    code,
		Data:    data,
		Message: "success",
	}
	c.JSON(code, res)
}
