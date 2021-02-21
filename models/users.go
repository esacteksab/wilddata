package models

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// ErrorNotFound is returned when a resource is not found in the DB
	ErrorNotFound = errors.New("models: resource not found")
)

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name  string
	Email string // `gorm: "not null;unique_index"`
}

func NewUserService() (*UserService, error) {
	// Openning file
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	// Error
	if err != nil {
		return nil, err
	}
	return &UserService{
		db: database,
	}, nil
}

// ByID will look up by the id provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 -
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrorNotFound
	default:
		return nil, err
	}
}

func (us *UserService) AutoMigrate() {
	us.db.AutoMigrate(&User{})
}
