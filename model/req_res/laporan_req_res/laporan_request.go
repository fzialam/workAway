package laporanreqres

type LaporanGetSPPDByIdRequest struct {
	UserId       int `validate:"required" json:"user_id"`
	SuratTugasId int `validate:"required" json:"surat_tugas_id"`
}

type UploadLaporanRequest struct {
	UserId         int    `validate:"required" json:"user_id"`
	SuratTugasId   int    `validate:"required" json:"surat_tugas_id"`
	DokLaporanName string `validate:"required" json:"dok_laporan_name"`
	DokLaporanPDF  string `validate:"required" json:"dok_laporan_pdf"`
}

type ApprovedLaporanRequest struct {
	Id      int    `validate:"required" json:"id"`
	UserId  int    `validate:"required" json:"user_id"`
	Status  string `validate:"required" json:"status"`
	Message string `validate:"required" json:"message"`
}
