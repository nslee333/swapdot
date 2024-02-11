package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var url Url

type Url struct {
	gorm.Model
	long_url  string
	short_url string
}

func ConnectToDatabase() (**gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Pacific",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	// fmt.Println(os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// fmt.Sprintf(dsn)

	if err != nil {
		// fmt.Printf("Error connecting to Database: %s, dsn: %s", err, dsn)

		fmt.Println(os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))

	}

	return &db, err
}

func main() {
	db, err := ConnectToDatabase()

	if err != nil {
		fmt.Println("")
	}

	// db
}

func GetUrl(db **gorm.DB) {

}

func InsertUrl(db **gorm.DB) {

}

func UpdateUrl(db **gorm.DB) {

}

func DeleteUrl(db **gorm.DB) {

}
