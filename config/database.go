package config

import (
	"app/model/domain"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// initializes the database and runs the initial migration
func Init() {
	ConnectDB()
	InitialMigration()
}

// retrieve the database configuration and open a connection to the database using GORM
func ConnectDB() {
	var err error

	configuration := GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configuration.DB_USERNAME,
		configuration.DB_PASSWORD,
		configuration.DB_HOST,
		configuration.DB_PORT,
		configuration.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect Database")
	}

	fmt.Println("Connected to Database")
}

func ConnectDBTest() *gorm.DB {
	var err error

	type Configuration struct {
		DB_USERNAME string
		DB_PASSWORD string
		DB_PORT     string
		DB_HOST     string
		DB_NAME     string
	}
	

	configuration := Configuration{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_PORT :     "3306",
		DB_HOST:     "127.0.0.1",
		DB_NAME:     "medisyncv",
	  }
	

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configuration.DB_USERNAME,
		configuration.DB_PASSWORD,
		configuration.DB_HOST,
		configuration.DB_PORT,
		configuration.DB_NAME,
	)

	fmt.Println(connectionString)
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect Database")
	}

	fmt.Println("Connected to Database")
	return DB
}

// runs an initial migration that creates tables in the database
func InitialMigration() {
	DB.AutoMigrate(
		&domain.User{},
		&domain.Category{},
		&domain.Medicine{},
		&domain.Schedule{},
		&domain.MedicalID{},
		&domain.MediChat{},
	)
}
