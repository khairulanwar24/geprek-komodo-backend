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
	IdOutlet       string `json:"id_outlet" form:"id_outlet" validate:"omitempty,uuid4"`
	JenisTransaksi string `json:"jenis_transaksi" form:"jenis_transaksi" validate:"omitempty,oneof=masuk keluar"`
}
