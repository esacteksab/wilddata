package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Assets struct
type Assets struct {
	gorm.Model
	ID   uint           //`gorm:"primaryKEY" json:"id"`
	Org  int            //`json:"org"`
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
