package entity

type Izin struct {
	Id                int    `json:"id"`
	SuratTugasId      int    `json:"surat_tugas_id"`
	UserId            int    `json:"user_id"`
	Status            string `json:"status"`
	CreateAt          string `json:"create_at"`
	StatusTTD         string `json:"status_ttd"`
	StatusTTDCreateAt string `json:"status_ttd_create_at"`
}

type ApprovedLaporanAk struct {
	Id           int    `json:"id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	UserId       int    `json:"user_id"`
	Status       string `json:"status"`
	CreateAt     string `json:"create_at"`
}

type ApprovedLaporanAngg struct {
	Id           int    `json:"id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	UserId       int    `json:"user_id"`
	Status       string `json:"status"`
	CreateAt     string `json:"create_at"`
}
