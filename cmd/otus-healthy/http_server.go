package main

import (
	"net/http"
	"time"
)

func NewServer(listen string) *http.Server {
	router := SetupRouter()

	return &http.Server{
		Handler:      router,
		Addr:         listen,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
