package entity

import "time"

type SuratTugas struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	JudulSurat string    `json:"judul_surat"`
	TglAwal    time.Time `json:"tgl_awal"`
	TglAkhir   time.Time `json:"tgl_akhir"`
	Create_at  time.Time `json:"create_at"`
	Status     int       `json:"status"`
}
