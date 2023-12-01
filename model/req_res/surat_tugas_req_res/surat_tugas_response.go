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

type SuratTugasJOINRincianResponse struct {
	Id               int                    `json:"id"`
	Tipe             int                    `json:"tipe"`
	UserId           int                    `json:"user_id"`
	LokasiTujuan     string                 `json:"lokasi_tujuan"`
	JenisProgram     string                 `json:"jenis_program"`
	DokumenName      string                 `json:"dokumen_name"`
	DokumenPDF       string                 `json:"dokumen_pdf"`
	DokPendukungName string                 `json:"dok_pendukung_name"`
	DokPendukungPdf  string                 `json:"dok_pendukung_pdf"`
	TglAwal          string                 `json:"tgl_awal"`
	TglAkhir         string                 `json:"tgl_akhir"`
	CreateAt         string                 `json:"create_at"`
	Rincian          entity.RincianAnggaran `json:"rincian"`
}

type SuratTugasJOINSPPDApprovedAnggaranResponse struct {
	Id               int                    `json:"id"`
	Tipe             int                    `json:"tipe"`
	UserId           int                    `json:"user_id"`
	LokasiTujuan     string                 `json:"lokasi_tujuan"`
	JenisProgram     string                 `json:"jenis_program"`
	DokumenName      string                 `json:"dokumen_name"`
	DokumenPDF       string                 `json:"dokumen_pdf"`
	DokPendukungName string                 `json:"dok_pendukung_name"`
	DokPendukungPdf  string                 `json:"dok_pendukung_pdf"`
	TglAwal          string                 `json:"tgl_awal"`
	TglAkhir         string                 `json:"tgl_akhir"`
	CreateAt         string                 `json:"create_at"`
	Status           string                 `json:"status"`
	Rincian          entity.RincianAnggaran `json:"rincian"`
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
	Message          string `json:"message"`
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
	UserNIP          string `json:"user_nip"`
	UserName         string `json:"user_name"`
	UserNoTelp       string `json:"user_no_telp"`
	UserEmail        string `json:"user_email"`
}

type SuratTugasJOINUserParticipanResponse struct {
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
	UserNIP          string                      `json:"user_nip"`
	UserName         string                      `json:"user_name"`
	UserNoTelp       string                      `json:"user_no_telp"`
	UserEmail        string                      `json:"user_email"`
	Participans      []entity.ParticipanJoinUser `json:"participans"`
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
	UserNIP          string                      `json:"user_nip"`
	UserName         string                      `json:"user_name"`
	UserNoTelp       string                      `json:"user_no_telp"`
	UserEmail        string                      `json:"user_email"`
	Participans      []entity.ParticipanJoinUser `json:"participans"`
}
type SuratTugasJOINDoubleApprovedUserParticipanResponse struct {
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
	OtherStatus      string                      `json:"other_status"`
	UserNIP          string                      `json:"user_nip"`
	UserName         string                      `json:"user_name"`
	UserNoTelp       string                      `json:"user_no_telp"`
	UserEmail        string                      `json:"user_email"`
	Participans      []entity.ParticipanJoinUser `json:"participans"`
}

type SuratTugasJOINApprovedUserParticipanLaporanResponse struct {
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
	UserNIP          string                      `json:"user_nip"`
	UserName         string                      `json:"user_name"`
	UserNoTelp       string                      `json:"user_no_telp"`
	UserEmail        string                      `json:"user_email"`
	Participans      []entity.ParticipanJoinUser `json:"participans"`
	Laporan          entity.Laporan              `json:"laporan"`
}

type SuratTugasJOINUserParticipanLaporanJOINApprovedResponse struct {
	Id               int                          `json:"id"`
	Tipe             int                          `json:"tipe"`
	UserId           int                          `json:"user_id"`
	LokasiTujuan     string                       `json:"lokasi_tujuan"`
	JenisProgram     string                       `json:"jenis_program"`
	DokumenName      string                       `json:"dokumen_name"`
	DokumenPDF       string                       `json:"dokumen_pdf"`
	DokPendukungName string                       `json:"dok_pendukung_name"`
	DokPendukungPdf  string                       `json:"dok_pendukung_pdf"`
	TglAwal          string                       `json:"tgl_awal"`
	TglAkhir         string                       `json:"tgl_akhir"`
	CreateAt         string                       `json:"create_at"`
	UserNIP          string                       `json:"user_nip"`
	UserName         string                       `json:"user_name"`
	UserNoTelp       string                       `json:"user_no_telp"`
	UserEmail        string                       `json:"user_email"`
	Participans      []entity.ParticipanJoinUser  `json:"participans"`
	Laporan          []entity.LaporanJoinApproved `json:"lpaoran"`
	Gambar           string                       `json:"gambar"`
	Lokasi           string                       `json:"lokasi"`
	Koordinat        string                       `json:"koordinat"`
}

type SuratTugasJOINUserFotoParticipanFotoLaporanStatusResponse struct {
	Id               int                             `json:"id"`
	Tipe             int                             `json:"tipe"`
	UserId           int                             `json:"user_id"`
	LokasiTujuan     string                          `json:"lokasi_tujuan"`
	JenisProgram     string                          `json:"jenis_program"`
	DokumenName      string                          `json:"dokumen_name"`
	DokumenPDF       string                          `json:"dokumen_pdf"`
	DokPendukungName string                          `json:"dok_pendukung_name"`
	DokPendukungPdf  string                          `json:"dok_pendukung_pdf"`
	TglAwal          string                          `json:"tgl_awal"`
	TglAkhir         string                          `json:"tgl_akhir"`
	CreateAt         string                          `json:"create_at"`
	UserNIP          string                          `json:"user_nip"`
	UserName         string                          `json:"user_name"`
	UserNoTelp       string                          `json:"user_no_telp"`
	UserEmail        string                          `json:"user_email"`
	Participans      []entity.ParticipanJoinUserFoto `json:"participans"`
	LaporanId        int                             `json:"laporan_id"`
	LaporanDokName   string                          `json:"laporan_dok_name"`
	LaporanDokPDF    string                          `json:"laporan_dok_pdf"`
	Status           string                          `json:"status"`
	NameGambar       string                          `json:"name_gambar"`
	Gambar           string                          `json:"gambar"`
	Lokasi           string                          `json:"lokasi"`
	Koordinat        string                          `json:"koordinat"`
}

type SuratTugasJOINLaporanApprovedResponse struct {
	Id               int                        `json:"id"`
	Tipe             int                        `json:"tipe"`
	UserId           int                        `json:"user_id"`
	LokasiTujuan     string                     `json:"lokasi_tujuan"`
	JenisProgram     string                     `json:"jenis_program"`
	DokumenName      string                     `json:"dokumen_name"`
	DokumenPDF       string                     `json:"dokumen_pdf"`
	DokPendukungName string                     `json:"dok_pendukung_name"`
	DokPendukungPdf  string                     `json:"dok_pendukung_pdf"`
	TglAwal          string                     `json:"tgl_awal"`
	TglAkhir         string                     `json:"tgl_akhir"`
	CreateAt         string                     `json:"create_at"`
	Laporan          entity.LaporanJoinApproved `json:"laporan"`
}

type SuratTugasJOINUserLaporanApprovedResponse struct {
	Id               int                        `json:"id"`
	Tipe             int                        `json:"tipe"`
	UserId           int                        `json:"user_id"`
	LokasiTujuan     string                     `json:"lokasi_tujuan"`
	JenisProgram     string                     `json:"jenis_program"`
	DokumenName      string                     `json:"dokumen_name"`
	DokumenPDF       string                     `json:"dokumen_pdf"`
	DokPendukungName string                     `json:"dok_pendukung_name"`
	DokPendukungPdf  string                     `json:"dok_pendukung_pdf"`
	TglAwal          string                     `json:"tgl_awal"`
	TglAkhir         string                     `json:"tgl_akhir"`
	CreateAt         string                     `json:"create_at"`
	Status           string                     `json:"status"`
	UserName         string                     `json:"user_name"`
	Laporan          entity.LaporanJoinApproved `json:"laporan"`
}
