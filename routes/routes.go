package routes

import (
	"net/http"
	"os"

	assets "github.com/esacteksab/wilddata/controllers/assets"
	auth "github.com/esacteksab/wilddata/controllers/auth"
	orgs "github.com/esacteksab/wilddata/controllers/orgs"
	"github.com/esacteksab/wilddata/controllers/token"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func StartGin() {

	port := os.Getenv("GOPORT")

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(sentrygin.New(sentrygin.Options{}))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"http://localhost:port"},
		AllowMethods:           []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:           append([]string{"content-type"}, supertokens.GetAllCORSHeaders()...),
		AllowCredentials:       true,
		ExposeHeaders:          []string{},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))
	router.Use(func(c *gin.Context){
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				c.Next()
			})).ServeHTTP(c.Writer, c.Request)
			c.Abort()
	})

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

	router.Run(":" + port)
}
