package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Artifacts struct
type Artifacts struct {
	gorm.Model
	ID   uint           //`gorm:"primaryKEY" json:"id"`
	Org  int            //`json:"org"`
	Name string         //`gorm:"not null" json:"name"`
	Tags datatypes.JSON // `json:"tags"`
}

func NewArtifactService() (*ArtifactService, error) {
	// Openning file
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	// Error
	if err != nil {
		return nil, err
	}
	return &ArtifactService{
		db: database,
	}, nil
}

type ArtifactService struct {
	db *gorm.DB
}

// APIV1GetArtifacts gets all artifacts
func (us *ArtifactService) APIV1GetArtifacts(c *gin.Context) {

	as, err := NewArtifactService()

	var artifacts []Artifacts
	c.BindJSON(&artifacts)
	err = as.db.Find(&artifacts).Error
	switch err {
	case nil:
		c.JSON(200, artifacts)
	case gorm.ErrRecordNotFound:
		c.JSON(404, gin.H{"not found": ErrorNotFound})
	default:
		c.JSON(404, gin.H{"error": err.Error})
	}

}

//// APIV1AddArtifact adds an artifact
//func APIV1AddArtifact(c *gin.Context) {
//
//	//db := InitDb()
//
//	var artifacts Artifacts
//	c.BindJSON(&artifacts)
//	fmt.Println(artifacts)
//	db.Create(&artifacts)
//	fmt.Println(artifacts)
//	c.JSON(201, gin.H{"success": artifacts})
//}
//
//// APIV1GetArtifact gets an individual artifact
//func APIV1GetArtifact(c *gin.Context) {
//
//	//db := InitDb()
//
//	var artifact Artifacts
//	id := c.Params.ByName("id")
//
//	sid, _ := strconv.Atoi(id)
//
//	db.Find(&artifact, "org = ?", sid)
//	fmt.Println(artifact)
//	c.JSON(200, artifact)
//}
//
//// APIV1UpdateArtifact updates an individual artifact
//func APIV1UpdateArtifact(c *gin.Context) {
//	id := c.Param("id")
//	c.JSON(200, gin.H{"method": "PUT", "id": id})
//}
//
//// APIV1DeleteArtifact deletes an individual artifact
//func APIV1DeleteArtifact(c *gin.Context) {
//	id := c.Param("id")
//	c.JSON(200, gin.H{"method": "DELETE", "id": id})
//}
//
