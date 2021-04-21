package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	httpPort := flag.String("http-port", "8000", "--http-port=8000")
	flag.Parse()

	if *httpPort == "" {
		logrus.Fatal("http-port is not specified")
		return
	}

	server := NewServer(fmt.Sprintf(":%s", *httpPort))

	logrus.Infof("HTTP server started and listening on: '%s' port", *httpPort)
	logrus.Fatal(server.ListenAndServe())
}
