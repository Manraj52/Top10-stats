package common

import (
	"os"
	"strings"
)

func GetGameTrackerServerIps() (multipleIpPorts []string) {
	gameTrackerIps := os.Getenv("GAME_TRACKER_SERVER_IP")
	ipPorts := strings.Split(gameTrackerIps, ";")

	for _, ips := range ipPorts {
		multipleIpPorts = append(multipleIpPorts, strings.TrimSpace(ips))
	}
	return
}

func GetServingDirectory() string {
	return os.Getenv("SERVER_NAME")
}

func GetListenPort() string {
	return os.Getenv("LISTEN_PORT")
}
