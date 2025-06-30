// controllers/finance_controller.go
package controllers

import (
	"ayam-geprek-backend/models"
	"ayam-geprek-backend/types"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Form input untuk transaksi
type CreateTransaksiForm struct {
	JenisTransaksi string  `json:"jenis_transaksi" form:"jenis_transaksi" validate:"required,oneof=masuk keluar"`
	Nominal        int     `json:"nominal" form:"nominal" validate:"required,min=0"`
	Keterangan     string  `json:"keterangan" form:"keterangan"`
	IdOutlet       string  `json:"id_outlet" form:"id_outlet" validate:"required,uuid4"`
	IdStokBahan    *string `json:"id_stok_bahan" form:"id_stok_bahan" validate:"omitempty,uuid4"`
	Jumlah         *int    `json:"jumlah" form:"jumlah" validate:"omitempty,min=0"`
}

func CreateTransaksi(c *fiber.Ctx) error {
	form := c.Locals("validatedForm").(*CreateTransaksiForm)

	idUser, ok := c.Locals("id_user").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(types.Response{
			Success: false,
			Message: "ID user tidak valid di token",
			Data:    nil,
		})
	}
	data := models.CreateTransaksi(
		form.JenisTransaksi,
		form.Nominal,
		form.Keterangan,
		form.IdOutlet,
		form.IdStokBahan,
		form.Jumlah,
		idUser,
	)

	return c.JSON(data)
}

func GetAllTransaksi(c *fiber.Ctx) error {
	form := c.Locals("validatedForm").(*types.GetData)
	filter := c.Locals("validatedForm2").(*types.GetDataTransaksi)

	var data types.Response

	// Jika filter global diisi, pakai pencarian global
	if form.Filter != "" {
		data = models.GetListTransaksiGlobal(form)
	} else {
		data = models.GetListTransaksiFilter(filter, form)
	}

	return c.JSON(data)
}
