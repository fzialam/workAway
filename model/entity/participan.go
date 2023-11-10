package entity

type Participan struct {
	Id           int   `json:"id"`
	UserId       []int `json:"user_id"`
	SuratTugasId int   `json:"surat_tugas_id"`
}

type ParticipanJoinUser struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	NIP          string `json:"nip"`
	Name         string `json:"name"`
	NoTelp       string `json:"no_telp"`
	Email        string `json:"email"`
}
