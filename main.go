package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nslee333/swapdot/database"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// db, err := database.ConnectToDatabase()

	// if err != nil {
	// 	fmt.Printf("Error connecting to database: %s", err)
	// }

	// db

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	// r := setupRouter()

	// r.Run(":3000")
	db, err := database.ConnectToDatabase()

	if err != nil {
		fmt.Println(db)
		return
	}

}
