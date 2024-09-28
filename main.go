package main

import (
	"irctc/database"
	"irctc/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.InitDatabase()
}

func main() {
	//fmt.Println("hello")

	router := gin.Default()

	routes.UserRoutes(router)
	routes.TrainRoutes(router)
	router.Run(":3000")
}
