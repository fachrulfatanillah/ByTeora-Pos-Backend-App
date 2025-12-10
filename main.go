package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/route"
)

func main() {
	// Load .env
	godotenv.Load()

	// Connect Database
	config.ConnectDB()

	// Run Migrations
	config.RunMigrations()

	// Start Gin Server
	r := gin.Default()
	route.RegisterRoutes(r)
	r.Run(":8080")
}