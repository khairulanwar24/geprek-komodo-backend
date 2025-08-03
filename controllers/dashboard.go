package controllers

import "github.com/gofiber/fiber/v2"

func IndexDashboard(c *fiber.Ctx) error {
	return c.Render("dashboard/index", fiber.Map{
		"PageTitle": "Dashboard",
		"Username":  c.Locals("username"),
	}, "layouts/main/layout")

}
