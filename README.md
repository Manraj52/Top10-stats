# Top10-stats
list Top 10 server stats from GameTracker

!!! Remember your gameTracker server must be verified !!!

    A) Edit before you build in main.go
    1. Open .prod-env & Add your gametracker server ip and port = ip:port
    2. (optional) Add serving directory path. ex: /server/sd
    3. (optional) Add listen port. ex: 9000, default: 80
    4. Edit in stats.gohtml "Ftag Europe Top10" to something else

.

    B) Install and Build
    1. Download & Install Go
    2. For linux server =
        GOOS=linux go build -o bin/stats main.go
    2. For Windows server = 
        GOOS=windows go build -o bin/stats.exe main.go
    2. For Mac server =
        GOOS=darwin go build -o bin/stats_mac main.go

    C) To run follow these steps =
    1. save bin/stats, .prod-env and stats.gohtml together anywhere in your instance
    2. open http port or tcp port <custom port>
    3. run this command = ./stats prod &
    4. Visit ip or ip:<custom port>

or

    B) Install
    1. Download & Install Go
    2. open http port or tcp port <custom port>
    3. Run this = go run main.go prod &
    4. Visit ip or ip:<custom port>

example = ip:4000/server/sd and ip:4000/api/server/sd <br>
var gameTrackerServerIp = "185.107.96.152:28960" <br>
var serverName 			= "server/sd" //optional <br>
var listenPort 			= "4000" //optional 




