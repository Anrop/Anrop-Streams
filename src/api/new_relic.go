package api

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent"
	newrelicGorilla "github.com/newrelic/go-agent/_integrations/nrgorilla/v1"
)

const appName = "Streams"

var (
	newRelicApp *newrelic.Application
)

func SetupNewRelic(licenseKey string) {
	config := newrelic.NewConfig(appName, licenseKey)

	var err error
	app, err := newrelic.NewApplication(config)
	newRelicApp = &app

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting New Relic: %q", err)
		os.Exit(1)
	}
}

func InstrumentRoutes(r *mux.Router) *mux.Router {
	if newRelicApp != nil {
		return newrelicGorilla.InstrumentRoutes(r, *newRelicApp)
	}

	return r
}
