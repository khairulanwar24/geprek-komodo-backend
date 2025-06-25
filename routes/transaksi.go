package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesFinance(app *fiber.App) {
	keuangan := app.Group("/transaksi", middlewares.JWTProtected())
	keuangan.Post("/", middlewares.ValidateForm(&controllers.CreateTransaksiForm{}), controllers.CreateTransaksi)
}
