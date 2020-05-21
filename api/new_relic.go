package api

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	newrelicGorilla "github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

const appName = "Streams"

var (
	newRelicApp *newrelic.Application
)

// SetupNewRelic performs initial configuration
func SetupNewRelic(licenseKey string) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
	)
	newRelicApp = app

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting New Relic: %q", err)
		os.Exit(1)
	}
}

// InstrumentRoutes adds NewRelic routes to mux.Router
func InstrumentRoutes(r *mux.Router) *mux.Router {
	if newRelicApp != nil {
		return newrelicGorilla.InstrumentRoutes(r, newRelicApp)
	}

	return r
}
