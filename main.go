package main

import (
	"github.com/labstack/echo"
	"github.com/thimunri/logtest/handlers"
	"github.com/thimunri/logtest/parser"
	"flag"
	"fmt"
	"os"
)

func main() {

	flagParser := flag.Bool("parser", false, "Parse log files")
	var logParser parser.LogParser
	logPath := os.Getenv("LOG_PATH")
	logParser.LogPath = logPath

	flag.Parse()

	if *flagParser {
		fmt.Println("Init log parser ...")
		err := logParser.Init()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Echo instance
	e := echo.New()

	var lhandler handlers.LogHandler
	lhandler.LogPath = logPath
	lhandler.GenerateMockUsers()

	// Routes
	e.GET("/server1", lhandler.LogAction)
	e.GET("/server2", lhandler.LogAction)
	e.GET("/server3", lhandler.LogAction)
	e.GET("/server4", lhandler.LogAction)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}


