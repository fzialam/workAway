package persetujuanreqres

import (
	"time"

	"github.com/fzialam/workAway/model/entity"
)

type PersetujuanResponse struct {
	Status  int    ` json:"status"`
	Message string `json:"message"`
}

type SuratTugasJOINApprovedUserResponse struct {
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
	CreateAt         string `json:"create_at"`
	Status           int    `json:"status"`
	UserNIP          string `json:"user_nip"`
	UserName         string `json:"user_name"`
	UserNoTelp       string `json:"user_no_telp"`
	UserEmail        string `json:"user_email"`
}

type SuratTugasJOINApprovedUserParticipanResponse struct {
	Id               int                         `json:"id"`
	UserId           int                         `json:"user_id"`
	LokasiTujuan     string                      `json:"lokasi_tujuan"`
	JenisProgram     int                         `json:"jenis_program"`
	DokumenName      string                      `json:"dokumen_name"`
	DokumenPDF       string                      `json:"dokumen_pdf"`
	DokPendukungName string                      `json:"dok_pendukung_name"`
	DokPendukungPdf  string                      `json:"dok_pendukung_pdf"`
	TglAwal          string                      `json:"tgl_awal"`
	TglAkhir         string                      `json:"tgl_akhir"`
	CreateAt         time.Time                   `json:"CreateAt"`
	Status           int                         `json:"status"`
	UserNIP          string                      `json:"user_nip"`
	UserName         string                      `json:"user_name"`
	UserNoTelp       string                      `json:"user_no_telp"`
	UserEmail        string                      `json:"user_email"`
	Participans      []entity.ParticipanJoinUser `json:"participans"`
}
