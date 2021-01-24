package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type route struct {
	ctx *gin.Context
}

func (r route) Bind(arg interface{}) {
	r.ctx.Bind(arg)
}

func (r route) Request() *http.Request {
	return r.ctx.Request
}

func (r route) Response(data interface{}) {
	r.ctx.JSON(
		http.StatusOK,
		data,
	)
}
