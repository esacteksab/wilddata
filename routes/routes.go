package routes

import (
	"net/http"
	"os"

	assets "github.com/esacteksab/wilddata/controllers/assets"
	auth "github.com/esacteksab/wilddata/controllers/auth"
	orgs "github.com/esacteksab/wilddata/controllers/orgs"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartGin() {

	port := os.Getenv("GOPORT")

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(sentrygin.New(sentrygin.Options{}))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
	}))

	apiV1 := router.Group("/v1")
	{

		apiV1.GET("assets", assets.APIV1GetAssets)

		apiV1.POST("assets", assets.APIV1AddAsset)

		apiV1.GET("a/:id", assets.APIV1GetAsset)

		apiV1.PUT("a/:id", assets.APIV1UpdateAsset)

		apiV1.DELETE("a/:id", assets.APIV1DeleteAsset)

		apiV1.GET("orgs", orgs.APIV1GetOrgs)

		apiV1.POST("orgs", orgs.APIV1AddOrg)

		apiV1.GET("o/:name", orgs.APIV1GetOrg)

		apiV1.PUT("o/:id", orgs.APIV1UpdateOrg)

		apiV1.DELETE("o/:id", orgs.APIV1DeleteOrg)

		apiV1.GET("o/:name/assets", orgs.APIV1GetOrgAssets)

		apiV1.GET("u/:name", auth.APIV1GetUser)

		apiV1.POST("users", auth.APIV1AddUser)

		apiV1.POST("auth/login", auth.APIV1Login)

		apiV1.POST("auth/logout", auth.APIV1Logout)

		apiV1.GET("users", auth.APIV1GetUsers)

	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":" + port)
}
