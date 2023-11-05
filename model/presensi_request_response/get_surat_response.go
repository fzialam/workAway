package presensirequestresponse

import "time"

type GetSuratForPresensiResponse struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	JudulSurat string    `json:"judul_surat"`
	TglAwal    time.Time `json:"tgl_awal"`
	TglAkhir   time.Time `json:"tgl_akhir"`
	Status     int       `json:"status"`
}
