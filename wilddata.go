package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esacteksab/wilddata/routes"
	"github.com/getsentry/sentry-go"
)

func main() {

	port := os.Getenv("GOPORT")
	sentryDSN := os.Getenv("SENTRY_DSN")
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	if sentryDSN == "" {
		log.Fatal("$SENTRY_DSN must be set")
	}

	if port == "" {
		log.Fatal("$GOPORT must be set")
	}

	routes.StartGin()
}
