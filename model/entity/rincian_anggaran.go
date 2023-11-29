package entity

type RincianAnggaran struct {
	Id           int    `json:"id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	DokName      string `json:"dok_name"`
	DokPDF       string `json:"dok_pdf"`
	CreateAt     string `json:"create_at"`
}
