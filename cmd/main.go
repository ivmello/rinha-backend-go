package main

import (
	"log"

	"rinha-backend-go/internal/infra/api"
	"rinha-backend-go/internal/infra/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	conn, err := database.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	api.Run(conn)
}
