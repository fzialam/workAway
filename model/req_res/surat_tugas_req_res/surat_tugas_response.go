package surattugasreqres

import (
	"github.com/fzialam/workAway/model/entity"
)

type SuratTugasResponse struct {
	Id               int    `json:"id"`
	Tipe             int    `json:"tipe"`
	UserId           int    `json:"user_id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     string `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
}

type SuratTugasJOINApprovedResponse struct {
	Id               int    `json:"id"`
	Tipe             int    `json:"tipe"`
	UserId           int    `json:"user_id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     string `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
	Status           string `json:"status"`
	StatusTTD        string `json:"status_ttd"`
}

type SuratTugasJOINApprovedLaporanResponse struct {
	Id               int    `json:"id"`
	Tipe             int    `json:"tipe"`
	UserId           int    `json:"user_id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     string `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
	CreateAt         string `json:"create_at"`
	StatusPimpinan   string `json:"status_pimpinan"`
	StatusKeuangan   string `json:"status_keuangan"`
}

type SuratTugasJOINApprovedLaporanDokumenResponse struct {
	Id               int    `json:"id"`
	Tipe             int    `json:"tipe"`
	UserId           int    `json:"user_id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     string `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
	CreateAt         string `json:"create_at"`
	LaporanAkName    string `json:"laporan_ak_name"`
	LaporanAkPDF     string `json:"laporan_ak_pdf"`
	LaporanAgName    string `json:"laporan_ag_name"`
	LaporanAgPDF     string `json:"laporan_ag_pdf"`
	StatusPimpinan   string `json:"status_pimpinan"`
	StatusKeuangan   string `json:"status_keuangan"`
}

type SuratTugasJOINApprovedUserResponse struct {
	Id               int    `json:"id"`
	Tipe             int    `json:"tipe"`
	UserId           int    `json:"user_id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     string `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
	CreateAt         string `json:"create_at"`
	Status           string `json:"status"`
	StatusTTD        string `json:"status_ttd"`
	UserNIP          string `json:"user_nip"`
	UserName         string `json:"user_name"`
	UserNoTelp       string `json:"user_no_telp"`
	UserEmail        string `json:"user_email"`
}

type SuratTugasJOINApprovedUserParticipanResponse struct {
	Id               int                         `json:"id"`
	Tipe             int                         `json:"tipe"`
	UserId           int                         `json:"user_id"`
	LokasiTujuan     string                      `json:"lokasi_tujuan"`
	JenisProgram     string                      `json:"jenis_program"`
	DokumenName      string                      `json:"dokumen_name"`
	DokumenPDF       string                      `json:"dokumen_pdf"`
	DokPendukungName string                      `json:"dok_pendukung_name"`
	DokPendukungPdf  string                      `json:"dok_pendukung_pdf"`
	TglAwal          string                      `json:"tgl_awal"`
	TglAkhir         string                      `json:"tgl_akhir"`
	CreateAt         string                      `json:"create_at"`
	Status           string                      `json:"status"`
	StatusTTD        string                      `json:"status_ttd"`
	UserNIP          string                      `json:"user_nip"`
	UserName         string                      `json:"user_name"`
	UserNoTelp       string                      `json:"user_no_telp"`
	UserEmail        string                      `json:"user_email"`
	Participans      []entity.ParticipanJoinUser `json:"participans"`
}

type SuratTugasJOINUserParticipanLaporanResponse struct {
	Id                   int                         `json:"id"`
	Tipe                 int                         `json:"tipe"`
	UserId               int                         `json:"user_id"`
	LokasiTujuan         string                      `json:"lokasi_tujuan"`
	JenisProgram         string                      `json:"jenis_program"`
	DokumenName          string                      `json:"dokumen_name"`
	DokumenPDF           string                      `json:"dokumen_pdf"`
	DokPendukungName     string                      `json:"dok_pendukung_name"`
	DokPendukungPdf      string                      `json:"dok_pendukung_pdf"`
	TglAwal              string                      `json:"tgl_awal"`
	TglAkhir             string                      `json:"tgl_akhir"`
	CreateAt             string                      `json:"create_at"`
	UserNIP              string                      `json:"user_nip"`
	UserName             string                      `json:"user_name"`
	UserNoTelp           string                      `json:"user_no_telp"`
	UserEmail            string                      `json:"user_email"`
	Participans          []entity.ParticipanJoinUser `json:"participans"`
	LaporanAktivitasName string                      `json:"laporan_aktivitas_name"`
	LaporanAktivitasPDF  string                      `json:"laporan_aktivitas_pdf"`
	LaporanAnggaranName  string                      `json:"laporan_anggaran_name"`
	LaporanAnggaranPDF   string                      `json:"laporan_anggaran_pdf"`
	Gambar               string                      `json:"gambar"`
	Lokasi               string                      `json:"lokasi"`
	Koordinat            string                      `json:"koordinat"`
}

type SuratTugasJOINUserFotoParticipanFotoResponse struct {
	Id                   int                             `json:"id"`
	Tipe                 int                             `json:"tipe"`
	UserId               int                             `json:"user_id"`
	LokasiTujuan         string                          `json:"lokasi_tujuan"`
	JenisProgram         string                          `json:"jenis_program"`
	DokumenName          string                          `json:"dokumen_name"`
	DokumenPDF           string                          `json:"dokumen_pdf"`
	DokPendukungName     string                          `json:"dok_pendukung_name"`
	DokPendukungPdf      string                          `json:"dok_pendukung_pdf"`
	TglAwal              string                          `json:"tgl_awal"`
	TglAkhir             string                          `json:"tgl_akhir"`
	CreateAt             string                          `json:"create_at"`
	UserNIP              string                          `json:"user_nip"`
	UserName             string                          `json:"user_name"`
	UserNoTelp           string                          `json:"user_no_telp"`
	UserEmail            string                          `json:"user_email"`
	Participans          []entity.ParticipanJoinUserFoto `json:"participans"`
	LaporanAktivitasName string                          `json:"laporan_aktivitas_name"`
	LaporanAktivitasPDF  string                          `json:"laporan_aktivitas_pdf"`
	LaporanAnggaranName  string                          `json:"laporan_anggaran_name"`
	LaporanAnggaranPDF   string                          `json:"laporan_anggaran_pdf"`
	NameGambar           string                          `json:"name_gambar"`
	Gambar               string                          `json:"gambar"`
	Lokasi               string                          `json:"lokasi"`
	Koordinat            string                          `json:"koordinat"`
}
