package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/types"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesKeuangan(app *fiber.App) {
	keuangan := app.Group("/transaksi", middlewares.JWTProtected())
	keuangan.Get("/",
		middlewares.ValidatedQueryAs("validatedForm", &types.GetData{}),
		middlewares.ValidatedQueryAs("validatedForm2", &types.GetDataTransaksi{}),
		controllers.GetAllTransaksi,
	)

	keuangan.Post("/", middlewares.ValidateForm(&controllers.CreateTransaksiForm{}), controllers.CreateTransaksi)
}
