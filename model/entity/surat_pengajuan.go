package entity

import "time"

type SuratPengajuan struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	TglAwal   time.Time `json:"tgl_awal"`
	TglAkhir  time.Time `json:"tgl_akhir"`
	Create_at time.Time `json:"create_at"`
}
