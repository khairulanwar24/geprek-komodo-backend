package middlewares

import "github.com/gofiber/fiber/v2"

func InjectMenu() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("MenuItems", []fiber.Map{
			{"title": "Dashboard", "url": "/dashboard", "icon": "bi bi-house"},
			{"title": "Master Kegiatan", "url": "/master-kegiatan", "icon": "bi bi-calendar"},
			{"title": "Absensi", "url": "/absensi", "icon": "bi bi-check2-circle"},
		})
		return c.Next()
	}
}
