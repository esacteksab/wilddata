package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Artifacts struct
type Artifacts struct {
	gorm.Model
	ID   uint           //`gorm:"primaryKEY" json:"id"`
	Org  int            //`json:"org"`
	Name string         //`gorm:"not null" json:"name"`
	Tags datatypes.JSON // `json:"tags"`
}

// InitDb intializes the Database
func InitDb() *gorm.DB {
	// Openning file
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	// Error
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	database.AutoMigrate(&Artifacts{})

	return database
}

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

	apiV1.GET("artifacts/:id", APIV1GetArtifact)

	apiV1.PUT("artifacts/:id", APIV1UpdateArtifact)

	apiV1.DELETE("artifacts/:id", APIV1DeleteArtifact)

	router.Run(":" + port)
}

// APIV1GetArtifacts gets all artifacts
func APIV1GetArtifacts(c *gin.Context) {

	db := InitDb()

	var artifacts []Artifacts
	fmt.Println(artifacts)
	db.Find(&artifacts)
	fmt.Println(artifacts)
	c.JSON(200, artifacts)
}

// APIV1AddArtifact adds an artifact
func APIV1AddArtifact(c *gin.Context) {

	db := InitDb()

	var artifacts Artifacts
	c.BindJSON(&artifacts)
	fmt.Println(artifacts)
	db.Create(&artifacts)
	fmt.Println(artifacts)
	c.JSON(201, gin.H{"success": artifacts})
}

// APIV1GetArtifact gets an individual artifact
func APIV1GetArtifact(c *gin.Context) {

	db := InitDb()

	var artifact Artifacts
	id := c.Params.ByName("id")

	sid, _ := strconv.Atoi(id)

	db.Find(&artifact, "org = ?", sid)
	fmt.Println(artifact)
	c.JSON(200, artifact)
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
