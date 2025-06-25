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
	JenisTransaksi string `json:"jenis_transaksi" form:"jenis_transaksi" validate:"required,oneof=masuk keluar"`
	Nominal        int    `json:"nominal" form:"nominal" validate:"required,min=0"`
	Keterangan     string `json:"keterangan" form:"keterangan"`
	IdOutlet       string `json:"id_outlet" form:"id_outlet" validate:"required,uuid4"`
	IdStokBahan    string `json:"id_stok_bahan" form:"id_stok_bahan" validate:"omitempty,uuid4"`
	Jumlah         *int   `json:"jumlah" form:"jumlah" validate:"omitempty,min=0"`
}

func CreateTransaksi(c *fiber.Ctx) error {
	form := c.Locals("validatedForm").(*CreateTransaksiForm)

	// Konversi UUID
	idOutlet, err := uuid.Parse(form.IdOutlet)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.Response{
			Success: false,
			Message: "ID Outlet tidak valid",
			Data:    nil,
		})
	}

	var idStokBahanPtr *uuid.UUID
	if form.IdStokBahan != "" {
		idStok, err := uuid.Parse(form.IdStokBahan)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(types.Response{
				Success: false,
				Message: "ID Stok Bahan tidak valid",
				Data:    nil,
			})
		}
		idStokBahanPtr = &idStok
	}

	resp := models.CreateTransaksi(
		form.JenisTransaksi,
		form.Keterangan,
		form.Nominal,
		&idOutlet,
		idStokBahanPtr,
		form.Jumlah,
	)

	return c.JSON(resp)
}
