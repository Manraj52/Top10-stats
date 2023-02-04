package stats

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"main/common"
	"strings"
)

var (
	ErrInvalidArguments = errors.New("Invalid Arguments ")
	ErrEmptyData        = errors.New("Fetch Data is empty ")
)

func getStats() (stats []place, err error) {
	gameTrackerServerIp := common.GetGameTrackerServerIps()

	if len(gameTrackerServerIp) < 0 {
		return nil, ErrInvalidArguments
	}

	var data []byte
	_, data, err = fasthttp.Get(data, fmt.Sprint("https://cache.gametracker.com/components/html0/?host=", gameTrackerServerIp[0], "&topPlayersHeight=135&showTopPlayers=1"))
	if err != nil {
		return
	}

	dataS := string(data)

	if len(dataS) < 0 {
		return nil, ErrEmptyData
	}

	d := dataS[strings.Index(dataS, "<b>Top 10 Players")+1:]

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
