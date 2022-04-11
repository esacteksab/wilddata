package assets

import (
	"fmt"

	"github.com/esacteksab/wilddata/models"
	"github.com/gin-gonic/gin"
)

// APIV1GetAssets gets all assets
func APIV1GetAllAssets(c *gin.Context) {

	db := models.InitDb()

	var assets []models.Assets
	db.Find(&assets)
	c.JSON(200, assets)
}

// APIV1AddAsset adds an asset
func APIV1AddAsset(c *gin.Context) {

	db := models.InitDb()

	var assets models.Assets
	c.BindJSON(&assets)
	fmt.Println(assets)
	db.Create(&assets)
	fmt.Println(assets)
	c.JSON(201, gin.H{"success": assets})
}

// APIV1GetAsset gets all assets with name "name"
func APIV1GetAsset(c *gin.Context) {

	db := models.InitDb()

	var assets []models.Assets
	db.Find(&assets)
	name := c.Params.ByName("name")



	// SELECT * from Assets where Org = `id`
	db.Find(&assets, "name = ?", name)
	fmt.Println(assets)
	c.JSON(200, assets)
}

// APIV1UpdateAsset updates an individual asset
// Assets are atomic. We don't want to allow updates
// func APIV1UpdateAsset(c *gin.Context) {
// 	id := c.Param("id")
// 	c.JSON(200, gin.H{"method": "PUT", "id": id})
// }

// APIV1DeleteAsset deletes an individual asset
func APIV1DeleteAsset(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"method": "DELETE", "id": id})
}