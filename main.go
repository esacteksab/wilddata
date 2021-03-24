package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Assets struct
type Assets struct {
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
	database.AutoMigrate(&Assets{})

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
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	apiV1 := router.Group("/v1")

	apiV1.GET("assets", APIV1GetAssets)

	apiV1.POST("assets", APIV1AddAsset)

	apiV1.GET("assets/:id", APIV1GetAsset)

	apiV1.PUT("assets/:id", APIV1UpdateAsset)

	apiV1.DELETE("assets/:id", APIV1DeleteAsset)

	router.Run(":" + port)
}

// APIV1GetAssets gets all assets
func APIV1GetAssets(c *gin.Context) {

	db := InitDb()

	var assets []Assets
	fmt.Println(assets)
	db.Find(&assets)
	fmt.Println(assets)
	c.JSON(200, assets)
}

// APIV1AddAsset adds an asset
func APIV1AddAsset(c *gin.Context) {

	db := InitDb()

	var assets Assets
	c.BindJSON(&assets)
	fmt.Println(assets)
	db.Create(&assets)
	fmt.Println(assets)
	c.JSON(201, gin.H{"success": assets})
}

// APIV1GetAsset gets an individual asset
func APIV1GetAsset(c *gin.Context) {

	db := InitDb()

	var asset Assets
	id := c.Params.ByName("id")

	sid, _ := strconv.Atoi(id)

	db.Find(&asset, "org = ?", sid)
	fmt.Println(asset)
	c.JSON(200, asset)
}

// APIV1UpdateAsset updates an individual asset
func APIV1UpdateAsset(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "PUT", "id": id})
}

// APIV1DeleteAsset deletes an individual asset
func APIV1DeleteAsset(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "DELETE", "id": id})
}
