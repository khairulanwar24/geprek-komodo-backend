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
	var stok []map[string]interface{}

	result := config.DB.Raw(`
		SELECT id_stok_bahan, nama_bahan, deskripsi, stok, satuan, kategori, updated_at
		FROM stok_bahan
		WHERE id_stok_bahan = ? AND status_data = true
		LIMIT 1
	`, id_stok_bahan).Scan(&stok)

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
	resp.Data = stok
	return resp
}

func UpdateStok(id_stok_bahan, nama_bahan, deskripsi string, stok int, satuan, kategori string) types.Response {
	var resp types.Response

	// ✅ Update dulu
	result := config.DB.Exec(`
		UPDATE stok_bahan SET 
			nama_bahan = ?, 
			deskripsi = ?, 
			stok = ?, 
			satuan = ?, 
			kategori = ?, 
			updated_at = now()
		WHERE id_stok_bahan = ? AND status_data = true
	`, nama_bahan, deskripsi, stok, satuan, kategori, id_stok_bahan)

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
		SELECT id_stok_bahan, nama_bahan, deskripsi, stok, satuan, kategori, updated_at
		FROM stok_bahan
		WHERE id_stok_bahan = ?
	`, id_stok_bahan).Scan(&updatedData)

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

func DeleteStok(id_stok_bahan string) types.Response {
	var resp types.Response

	// Eksekusi query untuk menghapus data master_aplikasi secara permanen
	result := config.DB.Exec(`DELETE FROM stok_bahan
								WHERE id_stok_bahan = ?`, id_stok_bahan)

	// Jika terjadi error saat eksekusi query
	if result.Error != nil {
		resp.Success = false
		resp.Message = "Gagal Menghapus Stok"
		resp.Data = nil
		return resp
	} else if result.RowsAffected == 0 {
		// Jika tidak ada data yang terpengaruh (id_stok_bahan tidak ditemukan)
		resp.Success = false
		resp.Message = "Stok tidak ditemukan"
		resp.Data = nil
		return resp
	}

	// Jika berhasil menghapus data
	resp.Success = true
	resp.Message = "Success"
	resp.Data = nil

	return resp
}
