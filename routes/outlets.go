package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/types"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesOutlets(app *fiber.App) {
	outlet := app.Group("/outlet", middlewares.JWTProtected())
	outlet.Get("/", middlewares.ValidatedParams2(&types.GetData{}), controllers.GetAllOutlets)
	outlet.Post("/", middlewares.ValidateForm(&controllers.CreateOutletForm{}), controllers.CreateOutlet)
}
