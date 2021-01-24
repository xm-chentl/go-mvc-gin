package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xm-chentl/go-mvc"
	ctxprop "github.com/xm-chentl/go-mvc/ctx-prop"
	"github.com/xm-chentl/go-mvc/enum"
)

type ginEx struct {
	handler     mvc.IHandler
	mode        enum.RouteMode
	routeFormat string
}

func (g *ginEx) SetHandle(handler mvc.IHandler) mvc.IMvc {
	g.handler = handler
	return g
}

func (g ginEx) Run(port int) {
	engine := gin.Default()
	engine.POST(g.route(), func(ctx *gin.Context) {
		// TODO... 可以以协程的方式优化，提高吞吐量
		g.handler.Exec(ctxprop.Context{
			enum.CTX:         route{ctx: ctx},
			enum.ServerName:  ctx.Param("server"),
			enum.ServiceName: ctx.Param("service"),
			enum.ActionName:  ctx.Param("action"),
			enum.RespFunc: func(data interface{}) {
				ctx.JSON(http.StatusOK, data)
			},
		})
	})
	fmt.Println("端口: ", port)
	engine.Run(
		fmt.Sprintf(":%d", port),
	)
}

func (g *ginEx) route() string {
	if g.routeFormat == "" {
		switch g.mode {
		case enum.ThreeMode:
			g.routeFormat = "/:server/:service/:action"
			break
		default:
			g.routeFormat = "/:service/:action"
		}
	}

	return g.routeFormat
}

// New 实例
func New() mvc.IMvc {
	return new(ginEx)
}

// NewMode 实例一个路由模式
func NewMode(mode enum.RouteMode) mvc.IMvc {
	return &ginEx{
		mode: mode,
	}
}
