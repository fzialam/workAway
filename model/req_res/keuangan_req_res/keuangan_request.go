package keuanganreqres

type UploadRincianAnggaranRequest struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	DokName      string `validate:"required" json:"dok_name"`
	DokPDF       string `validate:"required" json:"dok_pdf"`
}
