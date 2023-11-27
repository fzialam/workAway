package entity

type LaporanAnggaran struct {
	Id             int    `json:"id"`
	SuratTugasId   int    `json:"surat_tugas_id"`
	UserId         int    `json:"user_id"`
	DokLaporanName string `json:"dok_laporan_name"`
	DokLaporanPDF  string `json:"dok_laporan_pdf"`
	CreateAt       string `json:"create_at"`
}
