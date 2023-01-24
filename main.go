package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/valyala/fasthttp"
)

// Add your server ip and port here ip:port
// Add server Name (optional) in short with no spaces
// your server will goes to = ip:9000
// and api goes to = ip:9000/api/
// Add listen port (optional)
var gameTrackerServerIp = "185.107.96.152:28960"
var serverName = "" //optional
var listenPort = "" //optional

type stats struct {
	//UserProfile string
	Name  string `json:"name"`
	Score string `json:"score"`
}

type place struct {
	Stats    stats `json:"stats"`
	Position int   `json:"position"`
}

func main() {
	fmt.Printf("Server is Working\n\n")

	server := &fasthttp.Server{
		Handler: baseHandler,
	}

	port := ":"
	if strings.TrimSpace(listenPort) != "" {
		port += listenPort
	} else {
		port += "9000"
	}

	if err := server.ListenAndServe(port); err != nil {
		log.Println("ListenAndServe:", err)
		return
	}
}

func baseHandler(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()

	srvName := "/"
	if strings.TrimSpace(serverName) != "" {
		srvName += serverName
	}

	switch string(path) {
	case srvName:
		pageHandler(ctx)
	case "/api" + srvName:
		statsHandler(ctx)
	}
}

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

	top10 := struct {
		Top10 []place
	}{
		Top10: playerStats,
	}

	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)

	if err = tmpl.Execute(ctx, top10); err != nil {
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

func getStats() (stats []place, err error) {
	var data []byte
	_, data, err = fasthttp.Get(data, fmt.Sprint("https://cache.gametracker.com/components/html0/?host=", gameTrackerServerIp, "&topPlayersHeight=135&showTopPlayers=1"))
	if err != nil {
		return
	}

	dataS := string(data)
	d := dataS[strings.Index(dataS, "<b>Top 10 Players:"):]

	var z []string
	for i := 0; ; i++ {
		i = strings.Index(d, "https://www.gametracker.com/player/")

		if i <= 0 {
			break
		}

		z = append(z, d[i:])
		d = d[i+1:]
	}

	for i, s := range z {
		s = s[:strings.Index(s, "</div>\n\t\t\t<div class=\"item_float_clear")]

		stat := place{}
		stat.Position = i + 1
		//stat.Stats.UserProfile = s[:strings.Index(s, `"`)]
		stat.Stats.Name = strings.TrimSpace(s[strings.Index(s, `>`)+1 : strings.Index(s, `</a>`)])
		stat.Stats.Score = strings.TrimSpace(s[strings.Index(s, `scrollable_on_c03">`)+19:])

		// score := stat.Stats.Score
		// if strings.Contains(score, "k") {
		// 	parsedNumber, err := strconv.ParseFloat(score[:len(score)-1], 64)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	stat.Stats.Score = strconv.Itoa(int(parsedNumber * 1000))
		// }

		stats = append(stats, stat)
	}

	return
}
