package orgs

import (
	"fmt"

	"github.com/esacteksab/wilddata/models"
	"github.com/gin-gonic/gin"
)

// APIV1GetOrgs gets all assets
func APIV1GetOrgs(c *gin.Context) {

	db := models.InitDb()

	var assets []models.Orgs
	db.Find(&assets)
	c.JSON(200, assets)
}

// APIV1AddOrg adds an Org
func APIV1AddOrg(c *gin.Context) {

	db := models.InitDb()

	var orgs models.Orgs
	c.BindJSON(&orgs)

	db.Create(&orgs)

	c.JSON(201, gin.H{"success": orgs})
}

// APIV1GetOrg gets an individual Org
func APIV1GetOrg(c *gin.Context) {

	db := models.InitDb()

	var org models.Orgs
	name := c.Params.ByName("name")

	// SELECT * from Orgs where name = `name`
	db.Where("name = ?", name).Find(&org)
	c.JSON(200, org)
}

// APIV1GetOrgAssets gets an Org's Assets
func APIV1GetOrgAssets(c *gin.Context) {
	db := models.InitDb()

	var assets []models.Assets
	db.Find(&assets)
	name := c.Params.ByName("name")

	// SELECT * from Assets where Org = `id`
	db.Find(&assets, "org = ?", name).Find(&assets)
	fmt.Println(assets)
	c.JSON(200, assets)
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