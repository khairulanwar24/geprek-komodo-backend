package models

import (
	"ayam-geprek-backend/config"
	"ayam-geprek-backend/types"

	"github.com/google/uuid"
)

func CreateTransaksi(jenis, keterangan string, nominal int, idOutlet, idStokBahan *uuid.UUID, jumlah *int) types.Response {
	var resp types.Response

	// Validasi jenis
	if jenis != "masuk" && jenis != "keluar" {
		resp.Success = false
		resp.Message = "Jenis transaksi harus 'masuk' atau 'keluar'"
		return resp
	}

	// Insert transaksi
	query := `INSERT INTO transaksi_keuangan 
		(id_transaksi, jenis_transaksi, nominal, keterangan, id_outlet, waktu_transaksi, id_stok_bahan, jumlah, status_data)
		VALUES (?, ?, ?, ?, ?, now(), ?, ?, true)`

	newID := uuid.New()
	err := config.DB.Exec(query, newID, jenis, nominal, keterangan, idOutlet, idStokBahan, jumlah).Error
	if err != nil {
		resp.Success = false
		resp.Message = "Gagal menyimpan transaksi: " + err.Error()
		return resp
	}

	// Jika ada stok
	if idStokBahan != nil && jumlah != nil {
		stokUpdate := `UPDATE stok_bahan SET stok = stok + ? WHERE id_stok_bahan = ? AND status_data = true`
		if jenis == "keluar" {
			stokUpdate = `UPDATE stok_bahan SET stok = stok - ? WHERE id_stok_bahan = ? AND status_data = true`
		}

		err = config.DB.Exec(stokUpdate, *jumlah, *idStokBahan).Error
		if err != nil {
			resp.Success = false
			resp.Message = "Gagal update stok: " + err.Error()
			return resp
		}
	}

	resp.Success = true
	resp.Message = "Transaksi berhasil disimpan"
	resp.Data = nil
	// resp.Data = map[string]interface{}{
	// 	"id_transaksi":    newID,
	// 	"jenis_transaksi": jenis,
	// 	"nominal":         nominal,
	// 	"keterangan":      keterangan,
	// 	"id_outlet":       idOutlet,
	// 	"id_stok_bahan":   idStokBahan,
	// 	"jumlah":          jumlah,
	// 	"waktu_transaksi": time.Now(),
	// }
	return resp
}
