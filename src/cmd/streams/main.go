package main

import (
	"api"
	"db"
	"fmt"
	"net/http"
	"os"
	"twitch"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")
	newRelicLicenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	twitchClientID := os.Getenv("TWITCH_CLIENT_ID")

	if port == "" {
		port = "8080"
	}

	var err error

	err = db.Connect(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %q\n", err)
		os.Exit(1)
	}

	err = twitch.CreateSession(twitchClientID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing Twitch Session: %q\n", err)
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", api.StreamsHandler)

	if newRelicLicenseKey != "" {
		api.SetupNewRelic(newRelicLicenseKey)
		r = api.InstrumentRoutes(r)
	}

	var handler http.Handler
	handler = handlers.CORS()(r)
	handler = handlers.CompressHandler(handler)

	fmt.Fprintf(os.Stdout, "Server launching on port %s\n", port)

	// Bind to a port and pass our router in
	http.ListenAndServe(":"+port, handler)
}
