package token

import (
	"errors"
	"net/http"
	"time"

	"github.com/esacteksab/wilddata/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	EMail string `json:"email"`
}

func GenerateJWT(EMail string, Name string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		EMail: EMail,
		Name:  Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}


type TokenRequest struct {
	Name    string `json:"name"`
	Password string `json:"password"`
}
func GenerateToken(context *gin.Context) {
	db := models.InitDb()
	var request TokenRequest
	var user models.Users
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := db.Where("name = ?", request.Name).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err:= GenerateJWT(user.EMail, user.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}