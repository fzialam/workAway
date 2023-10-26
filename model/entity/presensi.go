package entity

import "time"

type Presensi struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	SuratTugasId int       `json:"surat_tugas_id"`
	Gambar       string    `json:"gambar"`
	Lokasi       string    `json:"lokasi"`
	Create_at    time.Time `json:"create_at"`
}