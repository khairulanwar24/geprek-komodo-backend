package main

import (
	"ayam-geprek-backend/config"
	"ayam-geprek-backend/models"
	"ayam-geprek-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Inisialisasi database & seeding user
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})
	models.SeedUsers()

	// Gunakan engine template HTML untuk render file dari ./views
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve file static dari folder public (assets, js, css)
	app.Static("/", "./public")

	// Middleware CORS (untuk frontend lain seperti React)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173", // React atau frontend lain
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,OPTIONS",
	}))

	// Handle OPTIONS (preflight)
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	// Setup semua route
	routes.SetupRoutes(app)

	// Jalankan server
	app.Listen(":3000")
}
