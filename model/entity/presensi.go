package entity

type Presensi struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	Gambar       string `json:"gambar"`
	Lokasi       string `json:"lokasi"`
	Koordinat    string `json:"koordinat"`
	CreateAt     string `json:"create_at"`
}
