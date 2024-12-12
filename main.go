package main

import (
	"echo-server/config"
	"echo-server/server"
	"flag"
	"fmt"
)

func parseArgs() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host on which to run the server to listen to incoming requests")
	flag.IntVar(&config.Port, "port", 8080, "port number on which to run the server to listen to incoming requests")

	flag.Parse()
}

func main() {

	parseArgs()

	fmt.Println("Running server on ", config.Host, ":", config.Port)

	server.RunServer()
}
