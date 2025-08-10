// controllers/menu_controller.go
package controllers

import (
	"ayam-geprek-backend/types"

	"github.com/gofiber/fiber/v2"
)

func GetMenu(c *fiber.Ctx) error {
	menu := map[string][]map[string]string{
		"Manajemen": {
			{"modul_path": "stock", "nama_modul": "Stok Bahan", "modul_icon": "ki-tablet"},
			{"modul_path": "laporan", "nama_modul": "Laporan", "modul_icon": "ki-clipboard"},
		},
	}

	return c.JSON(types.Response{
		Success: true,
		Message: "Menu loaded",
		Data: map[string]interface{}{
			"menu": menu,
		},
	})
}
