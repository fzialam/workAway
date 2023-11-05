package entity

import "time"

type Izin struct {
	Id           int       `json:"id"`
	SuratTugasId string    `json:"surat_tugas_id"`
	UserId       int       `json:"user_id"`
	Status       int       `json:"status"`
	Create_at    time.Time `json:"create_at"`
}
