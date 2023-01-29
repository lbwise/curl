package main

import (
	"os"
	"log"

	"github.com/lbwise/curl/cli"
	"github.com/lbwise/curl/client"
)

func main() {
	req, err := cli.ParseCmd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	client.SendRequest(req)
}