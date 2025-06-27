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
	outlet.Get("/:id_outlet", middlewares.ValidatedParams(&controllers.GetOutletByParams{}), controllers.GetOutletById)
	outlet.Put("/:id_outlet", middlewares.ValidatedParams(&controllers.UpdateOutletByParams{}), middlewares.ValidateForm(&controllers.UpdateOutletByForm{}), controllers.UpdateOutlet)
	outlet.Delete("/:id_outlet", middlewares.ValidatedParams(&controllers.DeleteOutletByParam{}), controllers.DeleteOutlet)

}
