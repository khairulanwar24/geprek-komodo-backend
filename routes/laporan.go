package routes

import (
	"ayam-geprek-backend/controllers"
	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/types"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesLaporan(app *fiber.App) {
	laporan := app.Group("/laporan", middlewares.JWTProtected())

	// 📊 Endpoint: GET /laporan/keuangan?start_date=...&end_date=...&id_outlet=...
	laporan.Get("/keuangan",
		middlewares.ValidatedQueryAs("validatedForm", &types.LaporanTransaksiQuery{}),
		controllers.LaporanKeuangan,
	)
}
