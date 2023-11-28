package entity

type Approved struct {
	Id                int    `json:"id"`
	SuratTugasId      int    `json:"surat_tugas_id"`
	UserId            int    `json:"user_id"`
	Status            string `json:"status"`
	Message           string `json:"message"`
	CreateAt          string `json:"create_at"`
	StatusTTD         string `json:"status_ttd"`
	MessageTTD        string `json:"message_ttd"`
	StatusTTDCreateAt string `json:"status_ttd_create_at"`
}

type ApprovedLaporan struct {
	Id        int    `json:"id"`
	LaporanId int    `json:"laporan_id"`
	UserId    int    `json:"user_id"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	CreateAt  string `json:"create_at"`
}
