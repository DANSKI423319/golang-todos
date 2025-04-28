package main

import (
	"golang-todos/database"
	"golang-todos/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var r *gin.Engine = gin.Default()

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

	routes.RegisterTodoRoutes(r)

	r.Run(":8080")
}
