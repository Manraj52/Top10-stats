package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"log"
	"main/common"
	"main/stats"
	"net"
	"os"
	"strings"
	"time"
)

type stdLogger struct{}

func (l *stdLogger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *stdLogger) Println(args ...interface{}) {
	log.Println(args...)
}

func (l *stdLogger) Print(args ...interface{}) {
	log.Print(args...)
}

func errorHandler(ctx *fasthttp.RequestCtx, err error) {
	ctx.Error("Something went wrong", fasthttp.StatusInternalServerError)
	log.Println(err)
}

func headerReceived(header *fasthttp.RequestHeader) fasthttp.RequestConfig {
	log.Println("Request received for", string(header.RequestURI()))

	return fasthttp.RequestConfig{
		ReadTimeout:        time.Second * 1,
		WriteTimeout:       time.Second * 1,
		MaxRequestBodySize: 1024,
	}
}

func continueHandler(header *fasthttp.RequestHeader) bool {
	log.Println("Received 100-continue request", header)
	fasthttp.StatusMessage(fasthttp.StatusContinue)
	return true
}

func connState(_ net.Conn, state fasthttp.ConnState) {
	fmt.Println("Connection state changed to ", state.String())
}

func main() {
	fmt.Printf("Server is Working\n\n")

	logger := &stdLogger{}

	environment := ""
	if len(os.Args) < 2 {
		log.Println("missing application environment method Usage: go run server.go {prod} ")
		os.Exit(1)
	} else {
		environment = os.Args[1]
	}

	if environment == "prod" {
		err := godotenv.Load(".prod-env")
		if err != nil {
			log.Println("failed to load .prod-env file, " + err.Error())
		}
	} else {
		logger.Println("Environment '" + environment + "' not available.\nUsage: go run server.go {prod} ")
		os.Exit(1)
	}

	server := &fasthttp.Server{
		Handler:               stats.BaseHandler,
		ErrorHandler:          errorHandler,
		HeaderReceived:        headerReceived,
		ContinueHandler:       continueHandler,
		Name:                  "ICS",
		Concurrency:           20,
		ReadBufferSize:        1024,
		WriteBufferSize:       128,
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
		IdleTimeout:           5 * time.Second,
		MaxConnsPerIP:         2,
		TCPKeepalivePeriod:    5 * time.Second,
		MaxRequestBodySize:    1 << 5, // !5 = 20
		DisableKeepalive:      true,
		ReduceMemoryUsage:     true,
		GetOnly:               true,
		SecureErrorLogMessage: true,
		//ConnState:             connState,
		Logger: logger,
	}

	listenPort := common.GetListenPort()
	port := ":"
	if strings.TrimSpace(listenPort) != "" {
		port += strings.TrimSpace(listenPort)
	} else {
		port += "80"
	}

	if err := server.ListenAndServe(port); err != nil {
		logger.Println("ListenAndServe:", err)
		return
	}
}
