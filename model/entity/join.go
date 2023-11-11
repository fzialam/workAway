package entity

import "time"

type SuratTugasJOINApprovedUserParticipan struct {
	Id               int                `json:"id"`
	UserId           int                `json:"user_id"`
	LokasiTujuan     string             `json:"lokasi_tujuan"`
	JenisProgram     int                `json:"jenis_program"`
	DokumenName      string             `json:"dokumen_name"`
	DokumenPDF       string             `json:"dokumen_pdf"`
	DokPendukungName string             `json:"dok_pendukung_name"`
	DokPendukungPdf  string             `json:"dok_pendukung_pdf"`
	TglAwal          string             `json:"tgl_awal"`
	TglAkhir         string             `json:"tgl_akhir"`
	CreateAt         time.Time          `json:"CreateAt"`
	Status           int                `json:"status"`
	UserNIP          string             `json:"user_nip"`
	UserName         string             `json:"user_name"`
	UserNoTelp       string             `json:"user_no_telp"`
	UserEmail        string             `json:"user_email"`
	Participans      ParticipanJoinUser `json:"participans"`
}
type SuratTugasJOINApprovedUser struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     int    `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
	CreateAt         string `json:"CreateAt"`
	Status           int    `json:"status"`
	UserNIP          string `json:"user_nip"`
	UserName         string `json:"user_name"`
	UserNoTelp       string `json:"user_no_telp"`
	UserEmail        string `json:"user_email"`
}

type SuratTugasJOINApproved struct {
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
	CreateAt         time.Time `json:"CreateAt"`
	Status           int       `json:"status"`
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
