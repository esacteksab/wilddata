package models

import (
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDb intializes the Database
func InitDb() *gorm.DB {
	// Openning file
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	// Error
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	database.Debug().AutoMigrate(&Assets{})
	database.Debug().AutoMigrate(&Orgs{})
	database.Debug().AutoMigrate(&Users{})

	return database
}

// Assets struct
type Assets struct {
	gorm.Model
	ID   uint           //`gorm:"primaryKEY" json:"id"`
	Org  string         //`json:"org"`
	Name string         //`gorm:"not null" json:"name"`
	Tags datatypes.JSON // `json:"tags"`
}

// Orgs Struct
type Orgs struct {
	gorm.Model
	//ID    uint   //`gorm:"primaryKEY" json: "id"`
	Name  string //`gorm:"not null" json:"name"`
	EMail string //`gorm:"unique_index:uidx_email" json:"email"`
}

// Users struct
type Users struct {
	gorm.Model
	Name  string //`gorm:"unique;not null" json:"name"`
	EMail     string //`gorm:"primaryKey;autoIncrement:false;unique_index;not null" json:"email"`
	CEmail    string //`gorm:"not null" json:"cemail"`
	VEMail    *bool  //`gorm:"default:true" json:"vemail,omitempty"`
	Password  string //`gorm:"not null" json:"password"`
	CPassword string //`gorm:"not null" json:"cpassword"`

}
