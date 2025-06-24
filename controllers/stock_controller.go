package controllers

import (
	"ayam-geprek-backend/models"
	"ayam-geprek-backend/types"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetListStok(c *fiber.Ctx) error {
	form := c.Locals("validatedForm").(*types.GetData)
	data := models.GetListStoks(form)
	return c.JSON(data)
}

type CreateStokForm struct {
	NamaBahan string `json:"nama_bahan" form:"nama_bahan" validate:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Stok      int    `json:"stok" form:"stok" validate:"required,min=0"`
	Satuan    string `json:"satuan" form:"satuan" validate:"required,oneof=kg liter pcs"`
	Kategori  string `json:"kategori" form:"kategori"`
}

func CreateStok(c *fiber.Ctx) error {
	// Ambil form yang sudah divalidasi
	form := c.Locals("validatedForm").(*CreateStokForm)

	// Panggil model untuk simpan ke DB
	data := models.CreateStok(
		form.NamaBahan,
		form.Deskripsi,
		form.Stok,
		form.Satuan,
		form.Kategori,
	)

	return c.JSON(data)
}

type GetStokByParams struct {
	Id_stok_bahan string `params:"id_stok_bahan" validate:"required,uuid4"`
}

func GetStokById(c *fiber.Ctx) error {
	form := c.Locals("validatedParams").(*GetStokByParams)

	data := models.GetStokById(form.Id_stok_bahan)

	return c.JSON(data)
}

type UpdateStokByParam struct {
	Id_stok_bahan string `json:"id_stok_bahan" params:"id_stok_bahan" validate:"required,uuid4"`
}

type UpdateStokByForm struct {
	Nama_Bahan string `json:"nama_bahan" form:"nama_bahan" validate:"required"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Stok       int    `json:"stok" form:"stok" validate:"required"`
	Satuan     string `json:"satuan" form:"satuan" validate:"required"`
	Kategori   string `json:"kategori" form:"kategori"`
}

func UpdateStok(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*UpdateStokByParam)
	fmt.Println("ID dari route:", params.Id_stok_bahan)
	id_stok_bahan := params.Id_stok_bahan

	form := c.Locals("validatedForm").(*UpdateStokByForm)

	nama_bahan := form.Nama_Bahan
	deskripsi := form.Deskripsi
	stok := form.Stok
	satuan := form.Satuan
	kategori := form.Kategori

	data := models.UpdateStok(id_stok_bahan, nama_bahan, deskripsi, stok, satuan, kategori)

	return c.JSON(data)

}

type DeleteStokByParam struct {
	Id_Stok_Bahan string `json:"id_stok_bahan" validate:"required,uuid4"`
}

func DeleteStok(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*DeleteStokByParam)
	id_stok_bahan := params.Id_Stok_Bahan

	data := models.DeleteStok(id_stok_bahan)

	return c.JSON(data)
}
