package persetujuanreqres

type PersetujuanRequest struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	Status       int    `validate:"required" json:"status"`
	CreateAt     string `validate:"required" json:"create_at"`
}