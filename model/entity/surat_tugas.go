package entity

import "time"

type SuratTugas struct {
	Id               int       `json:"id"`
	UserId           int       `json:"user_id"`
	LokasiTujuan     string    `json:"lokasi_tujuan"`
	JenisProgram     int       `json:"jenis_program"`
	DokumenName      string    `json:"dokumen_name"`
	DokumenPDF       string    `json:"dokumen_pdf"`
	DokPendukungName string    `json:"dok_pendukung_name"`
	DokPendukungPdf  string    `json:"dok_pendukung_pdf"`
	TglAwal          string    `json:"tgl_awal"`
	TglAkhir         string    `json:"tgl_akhir"`
	Create_at        time.Time `json:"create_at"`
}
