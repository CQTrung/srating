package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONRender struct{}

func (r *JSONRender) sendData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Status: "success",
		Data:   data,
	})
}

func (r *JSONRender) SendData(ctx *gin.Context, data interface{}) {
	r.sendData(ctx, data)
}

func (r *JSONRender) Success(ctx *gin.Context) {
	r.sendData(ctx, nil)
}

func (r *JSONRender) SendCustomData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
