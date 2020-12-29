package main

import (
	"fmt"
	"log"
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

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(sentrygin.New(sentrygin.Options{}))

	apiV1 := router.Group("/v1")

	apiV1.GET("artifacts", APIV1GetArtifacts)

	apiV1.POST("artifacts", APIV1AddArtifact)

	apiV1.GET("artifact/:id", APIV1GetArtifact)

	apiV1.PUT("artifact/:id", APIV1UpdateArtifact)

	apiV1.DELETE("artifact/:id", APIV1DeleteArtifact)
	router.Run(":" + port)
}

// APIV1GetArtifacts gets all artifacts
func APIV1GetArtifacts(c *gin.Context) {
	c.JSON(200, gin.H{"method": "GET"})
}

// APIV1AddArtifact adds an artifact
func APIV1AddArtifact(c *gin.Context) {
	c.JSON(200, gin.H{"method": "POST"})
}

// APIV1GetArtifact gets an individual artifact
func APIV1GetArtifact(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "GET", "id": id})
}

// APIV1UpdateArtifact updates an individual artifact
func APIV1UpdateArtifact(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "PUT", "id": id})
}

// APIV1DeleteArtifact deletes an individual artifact
func APIV1DeleteArtifact(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "DELETE", "id": id})
}
