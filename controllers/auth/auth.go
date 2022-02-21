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

	var users models.Users

	name := c.Params.ByName("name")

	db.Where("name = ?", name).Find(&users)
	c.JSON(200, users)
}

// APIV1AddUser adds a single user
func APIV1AddUser(c *gin.Context) {
	db := models.InitDb()

	var users models.Users

	c.BindJSON(&users)

	db.Create(&users)

	c.JSON(200, users)
}



// APIV1Login
func APIV1Login(c *gin.Context) {
	c.JSON(200, c)
}

// APIV1Logout
func APIV1Logout(c *gin.Context) {
	c.JSON(200, c)
}