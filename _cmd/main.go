package main

import (
	"fmt"
	"os"

	server "umlserver"
)

func main() {
	err := server.Listen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "run() error:\n%+v\n", err)
	}
}
