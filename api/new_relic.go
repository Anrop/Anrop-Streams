package api

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const appName = "Streams"

// SetupNewRelic sets up monitoring in NewRelic
func SetupNewRelic(licenseKey string, r *mux.Router) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting New Relic: %q", err)
		os.Exit(1)
	}

	r.Use(nrgorilla.Middleware(app))
}
