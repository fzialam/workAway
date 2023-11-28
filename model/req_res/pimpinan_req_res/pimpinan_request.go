package pimpinanreqres

type UploadSPPDRequest struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	DokName      string `validate:"required" json:"dok_name"`
	DokPDF       string `validate:"required" json:"dok_pdf"`
	Status       string `validate:"required" json:"status"`
	Message      string `validate:"required" json:"message"`
}
