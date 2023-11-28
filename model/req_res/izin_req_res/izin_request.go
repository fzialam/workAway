package izinreqres

type IzinRequest struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	Status       string `validate:"required" json:"status"`
	Message      string `json:"message"`
	StatusTTD    string `validate:"required" json:"status_ttd"`
	MessageTTD   string `json:"message_ttd"`
}
