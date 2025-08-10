package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesStock(app *fiber.App) {

	stok := app.Group("/stok-bahan", middlewares.JWTProtected())
	stok.Get("/", controllers.IndexStock)
	// stok.Get("/", middlewares.ValidatedParams2(&types.GetData{}), controllers.GetListStok)
	stok.Get("/add", controllers.AddStock)
	stok.Post("/", middlewares.ValidateForm(&controllers.CreateStokForm{}), controllers.CreateStok)
	// stok.Get("/:id_stok_bahan", middlewares.ValidatedParams(&controllers.GetStokByParams{}), controllers.GetStokById)
	// stok.Put("/:id_stok_bahan", middlewares.ValidatedParams(&controllers.UpdateStokByParam{}), middlewares.ValidateForm(&controllers.UpdateStokByForm{}), controllers.UpdateStok)
	// stok.Delete("/:id_stok_bahan", middlewares.ValidatedParams(&controllers.DeleteStokByParam{}), controllers.DeleteStok)
}
