package adminreqres

type IndexResponse struct {
	Permohonan       int `json:"permohonan"`
	Penugasan        int `json:"penugasan"`
	LaporanAktivitas int `json:"laporan_aktivitas"`
	LaporanAnggaran  int `json:"laporan_anggaran"`
	UserAk           int `json:"user_ak"`
	UserOff          int `json:"user_off"`
}
