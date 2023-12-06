package pimpinanreqres

type Permohonan struct {
	Belum    int `json:"belum"`
	Approved int `json:"approved"`
	Reject   int `json:"reject"`
}

type Penugasan struct {
	Belum int `json:"belum"`
	Sudah int `json:"sudah"`
}

type Laporan struct {
	Belum    int `json:"belum"`
	Approved int `json:"approved"`
	Reject   int `json:"reject"`
}

type IndexPimpinan struct {
	Permohonan Permohonan `json:"permohonan"`
	Penugasan  Penugasan  `json:"penugasan"`
	Laporan    Laporan    `json:"laporan"`
}
