package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	httpPort := os.Getenv("LISTEN")
	if httpPort == "" {
		logrus.Fatal("LISTEN env is not set")
		return
	}

	server := NewServer(fmt.Sprintf(":%s", httpPort))

	logrus.Infof("HTTP server started and listening on: '%s' port", httpPort)
	logrus.Fatal(server.ListenAndServe())
}
