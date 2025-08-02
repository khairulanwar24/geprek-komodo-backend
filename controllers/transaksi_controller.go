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

// controllers/transaksi_controller.go

func GetAllTransaksi(c *fiber.Ctx) error {
	form := c.Locals("global").(*types.GetData)
	filter := c.Locals("filter").(*types.GetDataTransaksi)

	// Jika global filter ada, gunakan pencarian umum
	if form.Filter != "" {
		return c.JSON(models.GetListTransaksiGlobal(form))
	}

	// Jika tidak, gunakan filter spesifik
	return c.JSON(models.GetListTransaksiFilter(filter, form))
}

type UpdateTransaksiParam struct {
	IdTransaksi string `params:"id_transaksi" validate:"required,uuid4"`
}

type UpdateTransaksiForm struct {
	JenisTransaksi string  `json:"jenis_transaksi" form:"jenis_transaksi" validate:"required,oneof=masuk keluar"`
	Nominal        int     `json:"nominal" form:"nominal" validate:"required,min=0"`
	Keterangan     string  `json:"keterangan" form:"keterangan"`
	IdOutlet       string  `json:"id_outlet" form:"id_outlet" validate:"required,uuid4"`
	IdStokBahan    *string `json:"id_stok_bahan" form:"id_stok_bahan" validate:"omitempty,uuid4"`
	Jumlah         *int    `json:"jumlah" form:"jumlah" validate:"omitempty,min=0"`
}

func UpdateTransaksi(c *fiber.Ctx) error {
	param := c.Locals("validatedParams").(*UpdateTransaksiParam)
	form := c.Locals("validatedForm").(*UpdateTransaksiForm)

	// Ambil user dari JWT
	idUser, ok := c.Locals("id_user").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(types.Response{
			Success: false,
			Message: "User tidak valid",
			Data:    nil,
		})
	}

	data := models.UpdateTransaksi(
		param.IdTransaksi,
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

type DeleteTransaksiParams struct {
	Id_Transaksi string `params:"id_transaksi" validate:"required,uuid4"`
}

func DeleteTransaksi(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*DeleteTransaksiParams)

	data := models.DeleteTransaksi(params.Id_Transaksi)
	return c.JSON(data)
}

type GetTransaksiByIdParam struct {
	Id_Transaksi string `params:"id_transaksi" validate:"required,uuid4"`
}

func GetTransaksiById(c *fiber.Ctx) error {
	param := c.Locals("validatedParams").(*GetTransaksiByIdParam)
	data := models.GetTransaksiById(param.Id_Transaksi)
	return c.JSON(data)
}

func LaporanKeuangan(c *fiber.Ctx) error {
	query := c.Locals("validatedForm").(*types.LaporanTransaksiQuery)
	data := models.GetLaporanKeuangan(query)
	return c.JSON(data)
}
