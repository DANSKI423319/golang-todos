package main

import (
	"golang-todos/database"
	"golang-todos/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	envError := godotenv.Load()
	if envError != nil {
		panic("Error loading .env file")
	}

	dbConnection, dbConnectionErr := database.ConnectToDb()
	if dbConnectionErr != nil {
		log.Fatalln(dbConnectionErr)
	}

	db, dbErr := dbConnection.DB()
	if dbErr != nil {
		log.Fatalln(dbErr)
	}

	dbErr = db.Ping()

	if dbErr != nil {
		log.Fatalln(dbErr)
	}

	defer db.Close()

	router := routes.SetupRouter()
	router.Run(":8080")
}
