package controllers

import (
	"ayam-geprek-backend/models"
	"ayam-geprek-backend/types"

	"github.com/gofiber/fiber/v2"
)

type CreateOutletForm struct {
	Nama   string `json:"nama_outlet" form:"nama_outlet" validate:"required"`
	Alamat string `json:"alamat" form:"alamat"`
}

func CreateOutlet(c *fiber.Ctx) error {
	form := c.Locals("validatedForm").(*CreateOutletForm)

	nama := form.Nama
	alamat := form.Alamat

	data := models.CreateOutlet(nama, alamat)

	return c.JSON(data)
}

func GetAllOutlets(c *fiber.Ctx) error {
	form := c.Locals("validatedForm").(*types.GetData)

	data := models.GetAllOutlets(form)

	return c.JSON(data)
}
