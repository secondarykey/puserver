package main

import (
	"flag"
	"fmt"
	"os"
	"umlserver/config"

	server "umlserver"
)

var jar string
var host string
var p int

func init() {
	flag.StringVar(&jar, "jar", "plantuml-nodot.jar", "jar file")
	flag.StringVar(&host, "host", "localhost", "listen server")
	flag.IntVar(&p, "p", 8080, "server port")
}

func main() {
	flag.Parse()
	err := server.Listen(config.Jar(jar), config.Port(p), config.Host(host))
	if err != nil {
		fmt.Fprintf(os.Stderr, "run() error:\n%+v\n", err)
	}
}
