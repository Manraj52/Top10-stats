# Top10-stats
list Top 10 server stats from Gametracker

!!! Remember your gametracker server must be verified !!!

A) Edit before you build in main.go
1. Add your server ip and port = ip:port
2. (optional) Add server Name in short with no spaces
3. (optional) Add listen port
4. Edit in stats.gohtml
    <div style="text-align: center;"><h2>Ftag Europe Top10</h2></div>

B) Install and Build
1. Download & Install Go
2. For linux server =
    GOOS=linux go build -o bin/stats main.go
2. For Windows server = 
    GOOS=windows go build -o bin/stats.exe main.go
2. For Mac server =
    GOOS=darwin go build -o bin/stats_mac main.go

C) To run follow these steps =
1. save bin/stats and stats.gohtml together anywhere in your instance
2. open tcp port 9000
3. run this command = ./stats &
4. Visit yourIp:9000

example = ip:4000/ics and ip:4000/api/ics
var gameTrackerServerIp = "185.107.96.152:28960"
var serverName 			= "ics" //optional 
var listenPort 			= "4000" //optional 




