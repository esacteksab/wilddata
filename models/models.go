package models

import (
	"log"
	"os"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDb intializes the Database
func InitDb() *gorm.DB {

	url := os.Getenv("DATABASE_URL")
	log.Println(url)
	dsn, _ := pq.ParseURL(url)
	log.Printf("DATABASE_URL:%v", dsn)

	// dsn := "host=db user=username password=password dbname=default_database port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

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
	Name  string //`gorm:"not null" json:"name"`
	DisplayName string //`gorm: "not null"`
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

func (user *Users) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *Users) CheckPassword(providedPassword string) error {
	log.Println(user.Password)
	log.Println(providedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
