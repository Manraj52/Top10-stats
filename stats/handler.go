package stats

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"html/template"
	"log"
)

func pageHandler(ctx *fasthttp.RequestCtx) {
	tmpl, err := template.ParseFiles("stats.gohtml")
	if err != nil {
		log.Println("template.ParseFiles:", err)
		return
	}

	playerStats, err := getStats()
	if err != nil {
		log.Println("getStats:", err)
		return
	}

	top := top10{
		Top10: playerStats,
	}

	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)

	if err = tmpl.Execute(ctx, top); err != nil {
		log.Println("tmpl.Execute:", err)
		return
	}
}

func statsHandler(ctx *fasthttp.RequestCtx) {
	playerStats, err := getStats()
	if err != nil {
		log.Println("getStats:", err)
		return
	}

	marshal, err := json.Marshal(playerStats)
	if err != nil {
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(marshal)
}
