package models

import (
	"ayam-geprek-backend/config"
	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/types"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateTransaksi(jenis string, nominal int, keterangan string, idOutlet string, idStokBahan *string, jumlah *int, createdBy uuid.UUID) types.Response {
	var resp types.Response

	// ✅ Konversi idOutlet string → uuid.UUID
	idOutletUUID, err := uuid.Parse(idOutlet)
	if err != nil {
		resp.Success = false
		resp.Message = "ID outlet tidak valid"
		return resp
	}

	// ✅ Konversi idStokBahan jika ada
	var idStokBahanUUID *uuid.UUID
	if idStokBahan != nil {
		parsed, err := uuid.Parse(*idStokBahan)
		if err != nil {
			resp.Success = false
			resp.Message = "ID stok bahan tidak valid"
			return resp
		}
		idStokBahanUUID = &parsed
	}

	// ✅ Simpan transaksi
	query := `
		INSERT INTO transaksi_keuangan (
			id_transaksi, jenis_transaksi, nominal, keterangan, id_outlet, waktu_transaksi,
			id_stok_bahan, jumlah, created_by, created_at, updated_at, status_data
		)
		VALUES (?, ?, ?, ?, ?, now(), ?, ?, ?, now(), now(), true)
	`
	newID := uuid.New()
	err = config.DB.Exec(query,
		newID,
		jenis,
		nominal,
		keterangan,
		idOutletUUID,
		idStokBahanUUID,
		jumlah,
		createdBy,
	).Error

	if err != nil {
		resp.Success = false
		resp.Message = "Gagal menyimpan transaksi: " + err.Error()
		return resp
	}

	// ✅ Optional update stok
	if idStokBahanUUID != nil && jumlah != nil {
		stokQuery := `UPDATE stok_bahan SET stok = stok + ? WHERE id_stok_bahan = ? AND status_data = true`
		if jenis == "keluar" {
			stokQuery = `UPDATE stok_bahan SET stok = stok - ? WHERE id_stok_bahan = ? AND status_data = true`
		}
		err = config.DB.Exec(stokQuery, *jumlah, *idStokBahanUUID).Error
		if err != nil {
			resp.Success = false
			resp.Message = "Gagal update stok: " + err.Error()
			return resp
		}
	}

	resp.Success = true
	resp.Message = "Transaksi berhasil disimpan"
	resp.Data = map[string]interface{}{
		"id_transaksi":    newID,
		"jenis_transaksi": jenis,
		"nominal":         nominal,
		"keterangan":      keterangan,
		"id_outlet":       idOutletUUID,
		"id_stok_bahan":   idStokBahanUUID,
		"jumlah":          jumlah,
		"created_by":      createdBy,
	}
	return resp
}

// models/finance.go

func GetListTransaksiGlobal(form *types.GetData) types.Response {
	var resp types.Response

	query := `
	SELECT 
		id_transaksi, jenis_transaksi, nominal, keterangan,
		id_outlet, id_stok_bahan, jumlah, created_by,
		waktu_transaksi, created_at, updated_at
	FROM transaksi_keuangan
	WHERE status_data = true
	`

	filter := ""
	if form.Filter != "" {
		val := "%" + strings.ToLower(form.Filter) + "%"
		filter += ` AND (
			LOWER(jenis_transaksi) LIKE '` + val + `' OR
			LOWER(keterangan) LIKE '` + val + `' OR
			CAST(nominal AS TEXT) LIKE '` + val + `'
		)`
	}

	data := middlewares.Datatables("", query, form.Order, filter, form.Limit, form.Offset)

	resp.Success = true
	resp.Message = "Data transaksi (filter global)"
	resp.Data = data
	return resp
}

func GetListTransaksiFilter(f *types.GetDataTransaksi, base *types.GetData) types.Response {
	var resp types.Response

	query := `
	SELECT 
		id_transaksi, jenis_transaksi, nominal, keterangan,
		id_outlet, id_stok_bahan, jumlah, created_by,
		waktu_transaksi, created_at, updated_at
	FROM transaksi_keuangan
	WHERE status_data = true
	`

	filter := ""
	if f.JenisTransaksi != "" {
		filter += ` AND jenis_transaksi = '` + f.JenisTransaksi + `'`
	}
	if f.IdOutlet != "" {
		filter += ` AND id_outlet = '` + f.IdOutlet + `'`
	}

	layout := "2006-01-02"
	if f.StartDate != "" && f.EndDate != "" {
		if _, err := time.Parse(layout, f.StartDate); err != nil {
			resp.Success = false
			resp.Message = "Format start_date tidak valid. Gunakan yyyy-mm-dd"
			return resp
		}
		if _, err := time.Parse(layout, f.EndDate); err != nil {
			resp.Success = false
			resp.Message = "Format end_date tidak valid. Gunakan yyyy-mm-dd"
			return resp
		}
		filter += ` AND DATE(waktu_transaksi) BETWEEN '` + f.StartDate + `' AND '` + f.EndDate + `'`
	}

	data := middlewares.Datatables(
		"", query, base.Order, filter, base.Limit, base.Offset,
	)

	resp.Success = true
	resp.Message = "Data transaksi (filter spesifik)"
	resp.Data = data
	return resp
}

func UpdateTransaksi(idTransaksi, jenis string, nominal int, keterangan, idOutlet string, idStokBahan *string, jumlah *int, updatedBy uuid.UUID) types.Response {
	var resp types.Response

	// Update data utama
	query := `
		UPDATE transaksi_keuangan
		SET jenis_transaksi = ?, nominal = ?, keterangan = ?, id_outlet = ?, 
			id_stok_bahan = ?, jumlah = ?, updated_at = now()
		WHERE id_transaksi = ? AND status_data = true
	`

	err := config.DB.Exec(query, jenis, nominal, keterangan, idOutlet, idStokBahan, jumlah, idTransaksi).Error
	if err != nil {
		resp.Success = false
		resp.Message = "Gagal update transaksi: " + err.Error()
		return resp
	}

	// Ambil data terbaru
	var data map[string]interface{}
	err = config.DB.Raw(`SELECT * FROM transaksi_keuangan WHERE id_transaksi = ?`, idTransaksi).Scan(&data).Error
	if err != nil {
		resp.Success = false
		resp.Message = "Update berhasil tapi gagal mengambil data"
		resp.Data = nil
		return resp
	}

	resp.Success = true
	resp.Message = "Transaksi berhasil diperbarui"
	resp.Data = data
	return resp
}

func DeleteTransaksi(id string) types.Response {
	var resp types.Response

	// Soft delete (status_data = false)
	result := config.DB.Exec(`
		UPDATE transaksi_keuangan 
		SET status_data = false, updated_at = now()
		WHERE id_transaksi = ? AND status_data = true
	`, id)

	if result.Error != nil {
		resp.Success = false
		resp.Message = "Gagal menghapus transaksi: " + result.Error.Error()
		resp.Data = nil
		return resp
	}

	if result.RowsAffected == 0 {
		resp.Success = false
		resp.Message = "Transaksi tidak ditemukan"
		resp.Data = nil
		return resp
	}

	resp.Success = true
	resp.Message = "Transaksi berhasil dihapus"
	resp.Data = nil
	return resp
}

// models/transaksi.go

func GetTransaksiById(id string) types.Response {
	var resp types.Response
	var data map[string]interface{}

	result := config.DB.Raw(`
		SELECT 
			t.id_transaksi,
			t.jenis_transaksi,
			t.nominal,
			t.keterangan,
			t.id_outlet,
			o.nama_outlet,
			t.id_stok_bahan,
			s.nama_bahan,
			t.jumlah,
			t.created_by,
			u.nama AS nama_user,
			t.waktu_transaksi,
			t.created_at,
			t.updated_at
		FROM transaksi_keuangan t
		LEFT JOIN outlet o ON t.id_outlet = o.id_outlet
		LEFT JOIN stok_bahan s ON t.id_stok_bahan = s.id_stok_bahan
		LEFT JOIN users u ON t.created_by = u.id_user
		WHERE t.id_transaksi = ? AND t.status_data = true
		LIMIT 1
	`, id).Scan(&data)

	if result.Error != nil {
		resp.Success = false
		resp.Message = "Gagal mengambil detail transaksi: " + result.Error.Error()
		resp.Data = nil
		return resp
	}

	if len(data) == 0 {
		resp.Success = false
		resp.Message = "Transaksi tidak ditemukan"
		resp.Data = nil
		return resp
	}

	resp.Success = true
	resp.Message = "Detail transaksi ditemukan"
	resp.Data = data
	return resp
}

func GetLaporanKeuangan(f *types.LaporanTransaksiQuery) types.Response {
	var resp types.Response

	filter := `WHERE status_data = true`
	if f.IdOutlet != "" {
		filter += ` AND id_outlet = '` + f.IdOutlet + `'`
	}
	if f.StartDate != "" && f.EndDate != "" {
		layout := "2006-01-02"
		if _, err := time.Parse(layout, f.StartDate); err != nil {
			resp.Success = false
			resp.Message = "Format start_date salah. Gunakan yyyy-mm-dd"
			return resp
		}
		if _, err := time.Parse(layout, f.EndDate); err != nil {
			resp.Success = false
			resp.Message = "Format end_date salah. Gunakan yyyy-mm-dd"
			return resp
		}
		filter += ` AND DATE(waktu_transaksi) BETWEEN '` + f.StartDate + `' AND '` + f.EndDate + `'`
	}

	query := `
	SELECT
		COALESCE(SUM(CASE WHEN jenis_transaksi = 'masuk' THEN nominal ELSE 0 END), 0) AS total_masuk,
		COALESCE(SUM(CASE WHEN jenis_transaksi = 'keluar' THEN nominal ELSE 0 END), 0) AS total_keluar,
		COALESCE(SUM(CASE WHEN jenis_transaksi = 'masuk' THEN nominal ELSE 0 END), 0) -
		COALESCE(SUM(CASE WHEN jenis_transaksi = 'keluar' THEN nominal ELSE 0 END), 0) AS saldo
	FROM transaksi_keuangan ` + filter

	result := make(map[string]interface{})
	err := config.DB.Raw(query).Scan(&result).Error
	if err != nil {
		resp.Success = false
		resp.Message = "Gagal mengambil data laporan: " + err.Error()
		resp.Data = nil
		return resp
	}

	resp.Success = true
	resp.Message = "Laporan berhasil diambil"
	resp.Data = result
	return resp
}
