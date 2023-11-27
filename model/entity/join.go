package entity

type SuratTugasJOINApprovedUserParticipan struct {
	Id               int                  `json:"id"`
	Tipe             int                  `json:"tipe"`
	UserId           int                  `json:"user_id"`
	LokasiTujuan     string               `json:"lokasi_tujuan"`
	JenisProgram     string               `json:"jenis_program"`
	DokumenName      string               `json:"dokumen_name"`
	DokumenPDF       string               `json:"dokumen_pdf"`
	DokPendukungName string               `json:"dok_pendukung_name"`
	DokPendukungPdf  string               `json:"dok_pendukung_pdf"`
	TglAwal          string               `json:"tgl_awal"`
	TglAkhir         string               `json:"tgl_akhir"`
	CreateAt         string               `json:"create_at"`
	Status           string               `json:"status"`
	StatusTTD        string               `json:"status_ttd"`
	UserNIP          string               `json:"user_nip"`
	UserName         string               `json:"user_name"`
	UserNoTelp       string               `json:"user_no_telp"`
	UserEmail        string               `json:"user_email"`
	Participans      []ParticipanJoinUser `json:"participans"`
}

type SuratTugasJOINUserParticipan struct {
	Id               int                  `json:"id"`
	Tipe             int                  `json:"tipe"`
	UserId           int                  `json:"user_id"`
	LokasiTujuan     string               `json:"lokasi_tujuan"`
	JenisProgram     string               `json:"jenis_program"`
	DokumenName      string               `json:"dokumen_name"`
	DokumenPDF       string               `json:"dokumen_pdf"`
	DokPendukungName string               `json:"dok_pendukung_name"`
	DokPendukungPdf  string               `json:"dok_pendukung_pdf"`
	TglAwal          string               `json:"tgl_awal"`
	TglAkhir         string               `json:"tgl_akhir"`
	CreateAt         string               `json:"create_at"`
	StatusTTD        string               `json:"status_ttd"`
	UserNIP          string               `json:"user_nip"`
	UserName         string               `json:"user_name"`
	UserNoTelp       string               `json:"user_no_telp"`
	UserEmail        string               `json:"user_email"`
	Participans      []ParticipanJoinUser `json:"participans"`
}

type SuratTugasJOINApprovedUser struct {
	Id                int    `json:"id"`
	Tipe              int    `json:"tipe"`
	UserId            int    `json:"user_id"`
	LokasiTujuan      string `json:"lokasi_tujuan"`
	JenisProgram      string `json:"jenis_program"`
	DokumenName       string `json:"dokumen_name"`
	DokumenPDF        string `json:"dokumen_pdf"`
	DokPendukungName  string `json:"dok_pendukung_name"`
	DokPendukungPdf   string `json:"dok_pendukung_pdf"`
	TglAwal           string `json:"tgl_awal"`
	TglAkhir          string `json:"tgl_akhir"`
	CreateAt          string `json:"create_at"`
	Status            string `json:"status"`
	StatusTTD         string `json:"status_ttd"`
	StatusTTDCreateAt string `json:"status_ttd_create_at"`
	UserNIP           string `json:"user_nip"`
	UserName          string `json:"user_name"`
	UserNoTelp        string `json:"user_no_telp"`
	UserEmail         string `json:"user_email"`
}

type SuratTugasJOINApproved struct {
	Id                int    `json:"id"`
	Tipe              int    `json:"tipe"`
	UserId            int    `json:"user_id"`
	LokasiTujuan      string `json:"lokasi_tujuan"`
	JenisProgram      string `json:"jenis_program"`
	DokumenName       string `json:"dokumen_name"`
	DokumenPDF        string `json:"dokumen_pdf"`
	DokPendukungName  string `json:"dok_pendukung_name"`
	DokPendukungPdf   string `json:"dok_pendukung_pdf"`
	TglAwal           string `json:"tgl_awal"`
	TglAkhir          string `json:"tgl_akhir"`
	CreateAt          string `json:"create_at"`
	Status            string `json:"status"`
	StatusTTD         string `json:"status_ttd"`
	StatusTTDCreateAt string `json:"status_ttd_create_at"`
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

type SuratTugasJOINApprovedLaporan struct {
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

type LaporanAktivitasAnggaran struct {
	SuratId          int    `json:"surat_id"`
	UserId           int    `json:"user_id"`
	DokAktivitasName string `json:"dok_aktivitas_name"`
	DokAktivitasPDF  string `json:"dok_aktivitas_pdf"`
	DokAnggaranName  string `json:"dok_anggaran_name"`
	DokAnggaranPDF   string `json:"dok_anggaran_pdf"`
}
