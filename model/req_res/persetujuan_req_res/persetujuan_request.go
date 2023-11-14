package persetujuanreqres

type PersetujuanRequest struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	Status       string `validate:"required" json:"status"`
	StatusTTD    string `validate:"required" json:"status_ttd"`
}
