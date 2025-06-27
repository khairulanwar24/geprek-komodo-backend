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
		INSERT INTO outlet (id_outlet, nama_outlet, alamat, status_data)
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
	sTable := `SELECT id_outlet, nama_outlet, alamat, created_at, updated_at FROM outlet WHERE status_data = true`

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

func GetOutletById(id_outlet string) types.Response {
	var resp types.Response
	var outlet []map[string]interface{}

	result := config.DB.Raw(`SELECT id_outlet, nama_outlet, alamat, 	
				created_at, updated_at FROM outlet
				WHERE id_outlet = ? AND status_data = true
				LIMIT 1
			`, id_outlet).Scan(&outlet)

	if result.Error != nil {
		resp.Success = false
		resp.Message = result.Error.Error()
		resp.Data = nil
	} else if result.RowsAffected == 0 {
		resp.Success = false
		resp.Message = "Outlet tidak ditemukan"
		resp.Data = nil
	}

	resp.Success = true
	resp.Message = "Outlet berhasil ditemukan"
	resp.Data = outlet

	return resp
}

func UpdateOutlet(id_outlet, nama_outlet, alamat string) types.Response {
	var resp types.Response

	// ✅ Update dulu
	result := config.DB.Exec(`
		UPDATE outlet SET 
			nama_outlet = ?, 
			alamat = ?, 

			updated_at = now()
		WHERE id_outlet = ? AND status_data = true
	`, nama_outlet, alamat, id_outlet)

	if result.Error != nil {
		resp.Data = nil
		resp.Message = "Gagal memperbarui data"
		resp.Success = false
		return resp
	} else if result.RowsAffected == 0 {
		resp.Data = nil
		resp.Message = "Data tidak ditemukan"
		resp.Success = false
		return resp
	}

	// ✅ Ambil data terbaru
	var updatedData map[string]interface{}
	result = config.DB.Raw(`
		SELECT id_outlet, nama_outlet, alamat, updated_at
		FROM outlet
		WHERE id_outlet = ?
	`, id_outlet).Scan(&updatedData)

	if result.Error != nil {
		resp.Data = nil
		resp.Message = "Gagal mengambil data setelah update"
		resp.Success = false
		return resp
	}

	resp.Data = updatedData
	resp.Message = "Data berhasil diperbarui"
	resp.Success = true
	return resp
}

func DeleteOutlet(id_outlet string) types.Response {
	var resp types.Response

	result := config.DB.Exec(`UPDATE outlet
								SET status_data = false, updated_at = now()
								WHERE id_outlet = ?`, id_outlet)

	if result.Error != nil {
		resp.Success = false
		resp.Message = "Gagal menghapus Data Outlet"
		resp.Data = nil
		return resp
	} else if result.RowsAffected == 0 {
		resp.Success = false
		resp.Message = "Data Outlet tidak ditemukan"
		resp.Data = nil
	}

	resp.Success = true
	resp.Message = "Success"
	resp.Data = nil

	return resp

}
