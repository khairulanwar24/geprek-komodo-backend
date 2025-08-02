package types

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetData struct {
	Limit  int    `json:"limit" form:"limit" validate:"required,numeric,oneof=10 25 50 100" default:"10"`
	Offset int    `json:"offset" form:"offset" validate:"numeric"`
	Order  string `json:"order" form:"order" validate:""`
	Filter string `json:"filter" form:"filter" validate:""`
	Params string `json:"params" form:"params" validate:""`
}

type GetDataTransaksi struct {
	IdOutlet       string `query:"id_outlet" form:"id_outlet" validate:"omitempty,uuid4"`
	JenisTransaksi string `query:"jenis_transaksi" form:"jenis_transaksi" validate:"omitempty,oneof=masuk keluar"`
	StartDate      string `json:"start_date" form:"start_date"`
	EndDate        string `json:"end_date" form:"end_date"`
}

type LaporanTransaksiQuery struct {
	IdOutlet  string `query:"id_outlet" validate:"omitempty,uuid4"`
	StartDate string `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   string `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
}
