package tureqres

type CreateSPPDRequest struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	DokumenName  string `validate:"required" json:"dokumen_name"`
	DokumenPDF   string `validate:"required" json:"dokumen_pdf"`
}
