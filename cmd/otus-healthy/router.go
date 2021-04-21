package main

import (
	"github.com/gorilla/mux"
	
	"github.com/d-ops-pro/healthy/monit"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().
		StrictSlash(true)
	{
		monit.ConfigureRouter(router)

	}

	return router
}
