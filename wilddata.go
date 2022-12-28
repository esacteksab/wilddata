package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esacteksab/wilddata/routes"
	"github.com/getsentry/sentry-go"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/dashboard/dashboardmodels"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {

	// The port Gin serves on. We set this explicitly to not conflict with Svelte
	port := os.Getenv("GOPORT")

	//
	// SuperTokens Setup and Config
	//
	stAPIBasePath := os.Getenv("ST_API_BASE_PATH")
	stWebsiteBasePath := os.Getenv("ST_WEBSITE_BASE_PATH")
	stConnectionURI := os.Getenv("ST_CONNECTION_URI")
	stAPIKey := os.Getenv("ST_API_KEY")
	stAppName := os.Getenv("ST_APP_NAME")
	stAPIDomain := os.Getenv("ST_API_DOMAIN")
	stWebSiteDomain := os.Getenv("ST_WEBSITE_DOMAIN")

	apiBasePath := stAPIBasePath
	websiteBasePath := stWebsiteBasePath
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
	  		ConnectionURI: stConnectionURI,
			APIKey: stAPIKey,
			},
		AppInfo: supertokens.AppInfo{
			AppName: stAppName,
			APIDomain: stAPIDomain,
			WebsiteDomain: stWebSiteDomain,
			APIBasePath: &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},

		RecipeList: []supertokens.Recipe{
			dashboard.Init(dashboardmodels.TypeInput{
				ApiKey: stAPIKey,
			}),
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})
    if err != nil {
		panic(err.Error())
	}

	//
	// Sentry Setup and config
	//
	sentryDSN := os.Getenv("SENTRY_DSN")
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Ensure Environment variables are set
	if sentryDSN == "" {
		log.Fatal("$SENTRY_DSN must be set")
	}

	if port == "" {
		log.Fatal("$GOPORT must be set")
	}

	routes.StartGin()
}
