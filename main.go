package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

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
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(sentrygin.New(sentrygin.Options{}))
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/davelist", func(c *gin.Context) {
		davelist := []string{"We", "are", "going", "to", "rock", "this", "shit."}
		c.HTML(http.StatusOK, "davelist.tmpl.html", gin.H{"list": davelist})
	})

	router.Run(":" + port)
}
