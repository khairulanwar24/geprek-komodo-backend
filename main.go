package main

import (
	"ayam-geprek-backend/config"
	"ayam-geprek-backend/models"
	"ayam-geprek-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})
	models.SeedUsers()

	app := fiber.New()

	routes.SetupRoutes(app)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	app.Listen(":3000")
}
