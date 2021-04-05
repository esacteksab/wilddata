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
	"gitlab.com/esacteksab/wtfizit/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDb intializes the Database
func InitDb() *gorm.DB {
	// Openning file
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	// Error
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	database.AutoMigrate(&models.Assets{})
	database.AutoMigrate(&models.Orgs{})

	return database
}

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

	apiV1.GET("orgs", APIV1GetOrgs)

	apiV1.POST("orgs", APIV1AddOrg)

	apiV1.GET("orgs/:id", APIV1GetOrg)

	apiV1.PUT("orgs/:id", APIV1UpdateOrg)

	apiV1.DELETE("orgs/:id", APIV1DeleteOrg)

	apiV1.GET("orgs:id/assets", APIV1GetOrgAssets)

	router.Run(":" + port)
}

// APIV1GetAssets gets all assets
func APIV1GetAssets(c *gin.Context) {

	db := InitDb()

	var assets []models.Assets
	db.Find(&assets)
	c.JSON(200, assets)
}

// APIV1AddAsset adds an asset
func APIV1AddAsset(c *gin.Context) {

	db := InitDb()

	var assets models.Assets
	c.BindJSON(&assets)
	fmt.Println(assets)
	db.Create(&assets)
	fmt.Println(assets)
	c.JSON(201, gin.H{"success": assets})
}

// APIV1GetAsset gets an individual asset
func APIV1GetAsset(c *gin.Context) {

	db := InitDb()

	var asset models.Assets
	id := c.Params.ByName("id")

	// id above is a string, we need an int
	sid, _ := strconv.Atoi(id)

	// SELECT * from Assets where Org = `id`
	db.Find(&asset, "id = ?", sid)
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

// APIV1GetOrgs gets all assets
func APIV1GetOrgs(c *gin.Context) {

	db := InitDb()

	var assets []models.Orgs
	db.Find(&assets)
	c.JSON(200, assets)
}

// APIV1AddOrg adds an Org
func APIV1AddOrg(c *gin.Context) {

	db := InitDb()

	var orgs models.Orgs
	c.BindJSON(&orgs)

	db.Create(&orgs)

	c.JSON(201, gin.H{"success": orgs})
}

// APIV1GetOrg gets an individual Org
func APIV1GetOrg(c *gin.Context) {

	db := InitDb()

	var org models.Orgs
	id := c.Params.ByName("id")

	// SELECT * from Orgs where ID = `id`
	db.Find(&org, id)
	c.JSON(200, org)
}

// APIV1GetOrgAssets gets an Org's Assets
func APIV1GetOrgAssets(c *gin.Context) {
	db := InitDb()

	var asset models.Assets
	id := c.Params.ByName("id")

	// id above is a string, we need an int
	sid, _ := strconv.Atoi(id)

	// SELECT * from Assets where Org = `id`
	db.Find(&asset, "org = ?", sid)
	fmt.Println(asset)
	c.JSON(200, asset)
}

// APIV1UpdateOrg updates an individual org
func APIV1UpdateOrg(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "PUT", "id": id})
}

// APIV1DeleteOrg deletes an individual org
func APIV1DeleteOrg(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "DELETE", "id": id})
}
