package controllers

import (
	"ayam-geprek-backend/models"
	"ayam-geprek-backend/types"
	"fmt"

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

type GetOutletByParams struct {
	Id_outlet string `params:"id_outlet" validate:"required,uuid4"`
}

func GetOutletById(c *fiber.Ctx) error {
	form := c.Locals("validatedParams").(*GetOutletByParams)

	data := models.GetOutletById(form.Id_outlet)

	return c.JSON(data)
}

type UpdateOutletByParams struct {
	Id_outlet string `json:"id_outlet" params:"id_outlet" validate:"required,uuid4"`
}

type UpdateOutletByForm struct {
	Nama_Outlet string `json:"nama_outlet" form:"nama_outlet" validate:"required"`
	Alamat      string `json:"alamat" form:"alamat"`
}

func UpdateOutlet(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*UpdateOutletByParams)
	fmt.Println("ID dari route:", params.Id_outlet)
	id_outlet := params.Id_outlet

	form := c.Locals("validatedForm").(*UpdateOutletByForm)

	nama_outlet := form.Nama_Outlet
	alamat := form.Alamat

	data := models.UpdateOutlet(id_outlet, nama_outlet, alamat)

	return c.JSON(data)

}

type DeleteOutletByParam struct {
	Id_outlet string `params:"id_outlet" validate:"required,uuid4"`
}

func DeleteOutlet(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*DeleteOutletByParam)
	id_outlet := params.Id_outlet

	data := models.DeleteOutlet(id_outlet)

	return c.JSON(data)
}
