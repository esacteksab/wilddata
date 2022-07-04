package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/esacteksab/wilddata/models"
	"github.com/gin-contrib/sessions"
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

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Hash the Password
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Hash the confirmation Password
	if err := user.HashCPassword(user.CPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Ensure Password and Confirmation Password (CPassword) match
	if errors.Is(user.CheckPassword(user.Password), user.CheckCPassword(user.CPassword)) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "passwords do not match"})
		c.Abort()
		return

	} else {
		u := db.Create(&user)
		if u.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": u.Error.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, gin.H{"userID": user.ID, "email": user.EMail, "name": user.Name})
	}

}

// APIV1Login
func APIV1Login(c *gin.Context) {

	db := models.InitDb()

	var user models.Users

	if c.ShouldBind(&user) == nil {
		// Map
		// db.Where(map[string]interface{}{"name": "user1", "password": "password1"}).Find(&users)
		// SELECT * FROM users WHERE name = "user1" AND password = "password1";
		if result := db.Where(map[string]interface{}{"name": user.Name, "password": user.Password}).Find(&user); result.Error != nil {
			c.String(200, "I am here")
		} else {
			if result.RowsAffected == 0 {
				log.Println(result.RowsAffected)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Username or Password"})
			} else {
				log.Println(result.RowsAffected)
				log.Println(map[string]interface{}{"message": user.ID})
				log.Println("SName: %s", user.Name)
				log.Println("SPasswd: %s", user.Password)
				c.JSON(http.StatusOK, gin.H{"message": "I found that user yo!"})

			}
		}

	} else {
		log.Println("You really fucked up! How did you get here?!")

	}
}

// APIV1Logout
func APIV1Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Signed out..."})
}

//AuthRequired ensures a session exists to access the endpoint
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("name")
	if sessionID == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "not authed",
		})
		c.Abort()
	}
}

// Ping will return Pong if a session exists
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}