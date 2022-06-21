package routes

import (
	"net/http"
	"os"

	assets "github.com/esacteksab/wilddata/controllers/assets"
	auth "github.com/esacteksab/wilddata/controllers/auth"
	orgs "github.com/esacteksab/wilddata/controllers/orgs"
	"github.com/esacteksab/wilddata/controllers/token"
	"github.com/esacteksab/wilddata/middlewares"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var (
	//RedisHost ...
	RedisHost = os.Getenv("REDIS_HOST")
	//RedisPort ...
	RedisPort = os.Getenv("REDIS_PORT")
)

func StartGin() {

	port := os.Getenv("GOPORT")

	store, _ := redis.NewStore(32, "tcp", RedisHost+":"+RedisPort, "", []byte("secret"))


	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(sentrygin.New(sentrygin.Options{}))
	router.Use(cors.Default())
	router.Use(sessions.Sessions("mysession", store))

	apiV1 := router.Group("/v1")
	{

		apiV1.GET("assets", assets.APIV1GetAllAssets)

		apiV1.POST("assets", assets.APIV1AddAsset)

		// Get all assets with the name "name"
		apiV1.GET("a/:name", assets.APIV1GetAsset)

		// assets are atomic. We don't want to allow updates
		//apiV1.PUT("a/:name", assets.APIV1UpdateAsset)

		apiV1.DELETE("a/:id", assets.APIV1DeleteAsset)

		apiV1.GET("orgs", orgs.APIV1GetOrgs)

		apiV1.POST("orgs", orgs.APIV1AddOrg)

		apiV1.GET("o/:name", orgs.APIV1GetOrg)

		apiV1.PUT("o/:id", orgs.APIV1UpdateOrg)

		apiV1.DELETE("o/:id", orgs.APIV1DeleteOrg)

		apiV1.GET("o/:name/assets", orgs.APIV1GetOrgAssets)

		apiV1.GET("u/:name", auth.APIV1GetUser)

		apiV1.POST("users", auth.APIV1AddUser)

		apiV1.POST("login", auth.APIV1Login)

		apiV1.POST("logout", auth.APIV1Logout)

		apiV1.GET("users", auth.APIV1GetUsers)

		apiV1.POST("token", token.GenerateToken)

	}

	authenticated := router.Group("/auth")
	authenticated.Use(middlewares.Auth())

	{
		authenticated.GET("ping", auth.Ping)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":" + port)
}
