package models

import "gorm.io/gorm"

// Orgs struct
type Orgs struct {
	gorm.Model
	Name  string //`json: "name"`
	EMail string // `json: "email"`
}
