package pegawaireqres

type Permohonan struct {
	Belum    int `json:"belum"`
	Approved int `json:"approved"`
	Reject   int `json:"reject"`
}

type Penugasan struct {
	Belum int `json:"belum"`
	Sudah int `json:"sudah"`
}

type Aktivitas struct {
	Belum    int `json:"belum"`
	Approved int `json:"approved"`
	Reject   int `json:"reject"`
}

type Anggaran struct {
	Belum    int `json:"belum"`
	Approved int `json:"approved"`
	Reject   int `json:"reject"`
}

type IndexPegawai struct {
	Permohonan Permohonan `json:"permohonan"`
	Penugasan  Penugasan  `json:"penugasan"`
	Aktivitas  Aktivitas  `json:"aktivitas"`
	Anggaran   Anggaran   `json:"anggaran"`
}
