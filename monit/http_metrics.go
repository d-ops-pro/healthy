package monit

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func ConfigureRouter(r *mux.Router) {
	r.Methods(http.MethodGet).Path("/health").HandlerFunc(healthCheckHandler)
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	result := map[string]interface{}{
		"status": "OK",
	}

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		logrus.WithError(err).Errorf("unable to write health response")
	}
}
