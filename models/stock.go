package models

import (
	"ayam-geprek-backend/config"

	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/types"
	"strings"
)

func GetListStoks(form *types.GetData) types.Response {
	var resp types.Response

	sRecursive := ``
	sTable := ` SELECT id_stok_bahan 
											, nama_bahan
											, deskripsi
											, stok
											, satuan
											, kategori
											, updated_at
											FROM
											stok_bahan WHERE status_data = true`

	sFilter := ``
	if form.Filter != "" {
		form.Filter = "%" + strings.ToLower(form.Filter) + "%"
		sFilter = `and LOWER(nama_bahan) LIKE ` + "'" + form.Filter + "'" + ` OR LOWER(deskripsi) LIKE ` + "'" + form.Filter + "'" + ` OR LOWER(stok) LIKE ` + "'" + form.Filter + "'" + ` OR LOWER(kategoru) LIKE ` + "'" + form.Filter + "'"
	} else {
		sFilter = ``
	}

	stock := middlewares.Datatables(
		sRecursive, sTable, form.Order, sFilter, form.Limit, form.Offset)

	resp.Success = true
	resp.Message = "Success"
	resp.Data = stock

	return resp

}

func CreateStok(nama_bahan, deskripsi string, stok int, satuan, kategori string) types.Response {
	var resp types.Response

	result := config.DB.Exec(`INSERT INTO stok_bahan (
		nama_bahan, deskripsi, stok, satuan, kategori, updated_at, status_data
	) VALUES (?, ?, ?, ?, ?, NOW(), TRUE)`,
		nama_bahan, deskripsi, stok, satuan, kategori,
	)

	if result.Error != nil {
		resp.Success = false
		resp.Message = result.Error.Error()
		return resp
	}

	resp.Success = true
	resp.Message = "Data stok berhasil ditambahkan"
	return resp
}

func GetStokById(id_stok_bahan string) types.Response {
	var resp types.Response
	var stock []map[string]interface{}

	result := config.DB.Raw(`
		SELECT id_stok_bahan, nama_bahan, deskripsi, stok, satuan, kategori, updated_at
		FROM stok_bahan
		WHERE id_stok_bahan = ? AND status_data = true
		LIMIT 1
	`, id_stok_bahan).Scan(&stock)

	if result.Error != nil {
		resp.Success = false
		resp.Message = result.Error.Error()
		resp.Data = nil
		return resp
	} else if result.RowsAffected == 0 {
		resp.Success = false
		resp.Message = "Data stok tidak ditemukan"
		resp.Data = nil
		return resp
	}

	resp.Success = true
	resp.Message = "Data stok berhasil ditemukan"
	resp.Data = stock
	return resp
}
