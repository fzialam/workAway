package entity

type RincianAnggaran struct {
	Id           int    `json:"id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	DokName      string `json:"dok_name"`
	DokPDF       string `json:"dok_pdf"`
	CreateAt     string `json:"create_at"`
}

type FullAnggaran struct {
	Id        int    `json:"id"`
	RincianId int    `json:"rincian_id"`
	Status    string `json:"status"`
	CreateAt  string `json:"create_at"`
}
