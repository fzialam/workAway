package entity

type Participan struct {
	Id           int   `json:"id"`
	UserId       []int `json:"user_id"`
	SuratTugasId int   `json:"surat_tugas_id"`
}
