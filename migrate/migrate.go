package main

import (
	"fmt"
	"log"

	"project-5/GORM-JWT/initializers"
	"project-5/GORM-JWT/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
