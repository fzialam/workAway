package presensireqres

type PresensiFotoRequest struct {
	UserId       int    `validate:"required" json:"user_id"`
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	Gambar       string `validate:"required" json:"gambar"`
	Lokasi       string `validate:"required" json:"lokasi"`
}
