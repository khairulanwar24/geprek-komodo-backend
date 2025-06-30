package models

import (
	"ayam-geprek-backend/config"
	"ayam-geprek-backend/middlewares"
	"ayam-geprek-backend/types"
	"strings"

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

func GetListTransaksiGlobal(form *types.GetData) types.Response {
	var resp types.Response

	sRecursive := ``
	sTable := `
	SELECT 
		id_transaksi,
		jenis_transaksi,
		nominal,
		keterangan,
		id_outlet,
		id_stok_bahan,
		jumlah,
		created_by,
		waktu_transaksi,
		created_at,
		updated_at
	FROM transaksi_keuangan
	WHERE status_data = true
`

	sFilter := ``
	if form.Filter != "" {
		form.Filter = "%" + strings.ToLower(form.Filter) + "%"
		sFilter += ` AND (
			LOWER(jenis_transaksi) LIKE '` + form.Filter + `' OR
			LOWER(keterangan) LIKE '` + form.Filter + `' OR
			CAST(nominal AS TEXT) LIKE '` + form.Filter + `'
		)`
	}

	data := middlewares.Datatables(
		sRecursive,
		sTable,
		form.Order,
		sFilter,
		form.Limit,
		form.Offset,
	)

	resp.Success = true
	resp.Message = "Data transaksi (pencarian global)"
	resp.Data = data
	return resp
}

// 🔍 2. Fungsi untuk filter spesifik
func GetListTransaksiFilter(form *types.GetDataTransaksi, base *types.GetData) types.Response {
	var resp types.Response

	sRecursive := ``
	sTable := `
	SELECT 
		id_transaksi,
		jenis_transaksi,
		nominal,
		keterangan,
		id_outlet,
		id_stok_bahan,
		jumlah,
		created_by,
		waktu_transaksi,
		created_at,
		updated_at
	FROM transaksi_keuangan
	WHERE status_data = true
`

	sFilter := ``
	if form.IdOutlet != "" {
		sFilter += ` AND id_outlet = '` + form.IdOutlet + `'`
	}
	if form.JenisTransaksi != "" {
		sFilter += ` AND jenis_transaksi = '` + form.JenisTransaksi + `'`
	}

	data := middlewares.Datatables(
		sRecursive,
		sTable,
		base.Order,
		sFilter,
		base.Limit,
		base.Offset,
	)

	resp.Success = true
	resp.Message = "Data transaksi (filter spesifik)"
	resp.Data = data
	return resp
}
