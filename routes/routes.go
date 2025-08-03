package routes

import (
	"ayam-geprek-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Static assets (jika perlu, bisa ditambah sesuai struktur)
	app.Static("/assets", "./public/assets")
	app.Static("/views", "./views") // optional, kalau perlu expose (biasanya tidak)

	// Global locals / layout injection (mirip mentor)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("Titleapp", "Ayam Geprek Admin")
		return c.Next()
	})

	// Auth / login routes
	SetupRoutesAuth(app)
	app.Use(middlewares.InjectMenu())
	// Setelah login: dashboard dengan middleware
	DashboardRoutes(app)

	// 404 fallback
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(404).Render("layouts/error/404", fiber.Map{
			"message":  "Halaman tidak ditemukan.",
			"Titleapp": c.Locals("Titleapp"),
		})
	})
}
