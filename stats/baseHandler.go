package stats

import (
	"github.com/valyala/fasthttp"
	"main/common"
	"strings"
)

func BaseHandler(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()

	serverName := common.GetServingDirectory()

	srvName := "/"
	if strings.TrimSpace(serverName) != "" {
		srvName += strings.TrimSpace(serverName)
	}

	switch string(path) {
	case srvName:
		pageHandler(ctx)
	case "/api" + srvName:
		statsHandler(ctx)
	}
}
