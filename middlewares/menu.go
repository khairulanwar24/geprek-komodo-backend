package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func InjectMenu() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Menu flat tanpa kategori
		menu := []map[string]string{
			{"modul_path": "stok-bahan", "nama_modul": "Stok Bahan", "modul_icon": "ki-tablet"},
			{"modul_path": "outlet", "nama_modul": "Outlet", "modul_icon": "ki-shop"},
			{"modul_path": "transaksi", "nama_modul": "Transaksi", "modul_icon": "ki-credit-cart"},
			{"modul_path": "laporan", "nama_modul": "Laporan", "modul_icon": "ki-clipboard"},
		}

		c.Locals("menu", menu)
		return c.Next()
	}
}
