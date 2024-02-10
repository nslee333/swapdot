package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var url Url

type Url struct {
	gorm.Model
	long_url  string
	short_url string
}

func db() {
	dsn := ""

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to Database: %s", err)
	}

}
