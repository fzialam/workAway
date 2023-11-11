package entity

type Izin struct {
	Id           int    `json:"id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	UserId       int    `json:"user_id"`
	Status       int    `json:"status"`
	CreateAt     string `json:"create_at"`
}
