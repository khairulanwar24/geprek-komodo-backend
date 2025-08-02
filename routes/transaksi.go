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
		middlewares.ValidatedQueryAs("global", &types.GetData{}),
		middlewares.ValidatedQueryAs("filter", &types.GetDataTransaksi{}),
		controllers.GetAllTransaksi,
	)
	keuangan.Put("/:id_transaksi",
		middlewares.ValidatedParams(&controllers.UpdateTransaksiParam{}),
		middlewares.ValidateForm(&controllers.UpdateTransaksiForm{}),
		controllers.UpdateTransaksi,
	)

	keuangan.Post("/", middlewares.ValidateForm(&controllers.CreateTransaksiForm{}), controllers.CreateTransaksi)
	keuangan.Delete("/:id_transaksi",
		middlewares.ValidatedParams(&controllers.DeleteTransaksiParams{}),
		controllers.DeleteTransaksi)

	keuangan.Get("/:id_transaksi",
		middlewares.ValidatedParams(&controllers.GetTransaksiByIdParam{}),
		controllers.GetTransaksiById)

}
