package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/RyanA3/hearth-gallery/internal/routes"
	"github.com/RyanA3/hearth-gallery/internal/templates"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/RyanA3/hearth-gallery/internal/models"
)

func main() {
	// Initialize template cache (important for production)
	err := templates.InitTemplateCache("./ui/html")
	if err != nil {
		log.Fatal(err)
	}

	// DB setup
	db, err := gorm.Open(sqlite.Open("hearth.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	// Get the router
	router := routes.NewRouter()

	fmt.Println("Server listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

