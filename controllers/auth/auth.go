package auth

import (
	"github.com/esacteksab/wilddata/models"
	"github.com/gin-gonic/gin"
)

// APIV1GetUsers gets all users
func APIV1GetUsers(c *gin.Context) {
	db := models.InitDb()

	var users []models.Users
	db.Find(&users)
	c.JSON(200, users)
}

// APIV1GetUser gets a single user
func APIV1GetUser(c *gin.Context) {
	db := models.InitDb()

	var user models.Users

	name := c.Params.ByName("name")

	db.Where("name = ?", name).Find(&user)
	c.JSON(200, user)
}

// APIV1AddUser adds a single user
func APIV1AddUser(c *gin.Context) {
	db := models.InitDb()

	var user models.Users

	c.BindJSON(&user)

	db.Create(&user)

	c.JSON(200, user)
}

// APIV1UpdateUser updates a single user