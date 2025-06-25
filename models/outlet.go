package models

import (
	"ayam-geprek-backend/config"
	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/types"
	"strings"

	"github.com/google/uuid"
)

func CreateOutlet(nama, alamat string) types.Response {
	var resp types.Response

	id := uuid.New()

	query := `
		INSERT INTO outlets (id_outlet, nama_outlet, alamat, status_data)
		VALUES (?, ?, ?, true)
	`

	err := config.DB.Exec(query, id, nama, alamat).Error
	if err != nil {
		resp.Success = false
		resp.Message = "Gagal menyimpan outlet: " + err.Error()
		return resp
	}

	resp.Success = true
	resp.Message = "Outlet berhasil ditambahkan"
	resp.Data = map[string]interface{}{
		"id_outlet":   id,
		"nama_outlet": nama,
		"alamat":      alamat,
	}
	return resp
}

func GetAllOutlets(form *types.GetData) types.Response {
	var resp types.Response

	sRecursive := ``
	sTable := `SELECT id_outlet, nama_outlet, alamat, created_at, updated_at FROM outlets WHERE status_data = true`

	// Handle filter (LOWER LIKE)
	sFilter := ``
	if form.Filter != "" {
		form.Filter = "%" + strings.ToLower(form.Filter) + "%"
		sFilter = `AND (LOWER(nama_outlet) LIKE '` + form.Filter + `' OR LOWER(alamat) LIKE '` + form.Filter + `')`
	}

	// Query datatables
	outletData := middlewares.Datatables(
		sRecursive, sTable, form.Order, sFilter, form.Limit, form.Offset,
	)

	resp.Success = true
	resp.Message = "Success"
	resp.Data = outletData
	return resp
}
