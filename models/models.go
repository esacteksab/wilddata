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
	database.AutoMigrate(&Assets{})
	database.AutoMigrate(&Orgs{})

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
	EMail string //`gorm:"unique_index";not null" json:"email"`
}
