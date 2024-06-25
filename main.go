package main

import (
	"learn-go/internal/api"
	"learn-go/internal/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()

	defer db.DB.Close()
	router := api.NewRouter()
	log.Fatal(router.Run(":8080"))
}
