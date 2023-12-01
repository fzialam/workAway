package entity

// Surat Tugas
type SuratTugasJOINApproved struct {
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
}

type SuratTugasJOINRincian struct {
	Id               int             `json:"id"`
	Tipe             int             `json:"tipe"`
	UserId           int             `json:"user_id"`
	LokasiTujuan     string          `json:"lokasi_tujuan"`
	JenisProgram     string          `json:"jenis_program"`
	DokumenName      string          `json:"dokumen_name"`
	DokumenPDF       string          `json:"dokumen_pdf"`
	DokPendukungName string          `json:"dok_pendukung_name"`
	DokPendukungPdf  string          `json:"dok_pendukung_pdf"`
	TglAwal          string          `json:"tgl_awal"`
	TglAkhir         string          `json:"tgl_akhir"`
	CreateAt         string          `json:"create_at"`
	Rincian          RincianAnggaran `json:"rincian"`
}

type SuratTugasJOINUserLaporanApproved struct {
	Id               int                 `json:"id"`
	Tipe             int                 `json:"tipe"`
	UserId           int                 `json:"user_id"`
	LokasiTujuan     string              `json:"lokasi_tujuan"`
	JenisProgram     string              `json:"jenis_program"`
	DokumenName      string              `json:"dokumen_name"`
	DokumenPDF       string              `json:"dokumen_pdf"`
	DokPendukungName string              `json:"dok_pendukung_name"`
	DokPendukungPdf  string              `json:"dok_pendukung_pdf"`
	TglAwal          string              `json:"tgl_awal"`
	TglAkhir         string              `json:"tgl_akhir"`
	CreateAt         string              `json:"create_at"`
	Status           string              `json:"status"`
	UserName         string              `json:"user_name"`
	Laporan          LaporanJoinApproved `json:"laporan"`
}

type SuratTugasJOINUser struct {
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

type SuratTugasJOINUserFoto struct {
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
	UserNIP          string `json:"user_nip"`
	UserName         string `json:"user_name"`
	UserNoTelp       string `json:"user_no_telp"`
	UserEmail        string `json:"user_email"`
	UserNameGambar   string `json:"user_name_gambar"`
	UserGambar       string `json:"user_gambar"`
	UserLokasi       string `json:"user_lokasi"`
	UserKoordinat    string `json:"user_koordinat"`
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
	StatusTTDCreateAt string `json:"status_ttd_create_at"`
	UserNIP           string `json:"user_nip"`
	UserName          string `json:"user_name"`
	UserNoTelp        string `json:"user_no_telp"`
	UserEmail         string `json:"user_email"`
}

type SuratTugasJOINApprovedUserOtherId struct {
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
	OtherStatus      string `json:"other_status"`
	UserNIP          string `json:"user_nip"`
	UserName         string `json:"user_name"`
	UserNoTelp       string `json:"user_no_telp"`
	UserEmail        string `json:"user_email"`
	OtherId          int    `json:"other_id"`
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
	Message          string `json:"message"`
}

type SuratTugasJOINSPPDApprovedAnggaran struct {
	Id               int             `json:"id"`
	Tipe             int             `json:"tipe"`
	UserId           int             `json:"user_id"`
	LokasiTujuan     string          `json:"lokasi_tujuan"`
	JenisProgram     string          `json:"jenis_program"`
	DokumenName      string          `json:"dokumen_name"`
	DokumenPDF       string          `json:"dokumen_pdf"`
	DokPendukungName string          `json:"dok_pendukung_name"`
	DokPendukungPdf  string          `json:"dok_pendukung_pdf"`
	TglAwal          string          `json:"tgl_awal"`
	TglAkhir         string          `json:"tgl_akhir"`
	CreateAt         string          `json:"create_at"`
	Status           string          `json:"status"`
	Rincian          RincianAnggaran `json:"rincian"`
}

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
	UserNIP          string               `json:"user_nip"`
	UserName         string               `json:"user_name"`
	UserNoTelp       string               `json:"user_no_telp"`
	UserEmail        string               `json:"user_email"`
	Participans      []ParticipanJoinUser `json:"participans"`
}

type SuratTugasJOINDoubleApprovedUserParticipan struct {
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
	OtherStatus      string               `json:"other_status"`
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
	UserNIP          string               `json:"user_nip"`
	UserName         string               `json:"user_name"`
	UserNoTelp       string               `json:"user_no_telp"`
	UserEmail        string               `json:"user_email"`
	Participans      []ParticipanJoinUser `json:"participans"`
}

type SuratTugasJOINUserParticipanFoto struct {
	Id               int                      `json:"id"`
	Tipe             int                      `json:"tipe"`
	UserId           int                      `json:"user_id"`
	LokasiTujuan     string                   `json:"lokasi_tujuan"`
	JenisProgram     string                   `json:"jenis_program"`
	DokumenName      string                   `json:"dokumen_name"`
	DokumenPDF       string                   `json:"dokumen_pdf"`
	DokPendukungName string                   `json:"dok_pendukung_name"`
	DokPendukungPdf  string                   `json:"dok_pendukung_pdf"`
	TglAwal          string                   `json:"tgl_awal"`
	TglAkhir         string                   `json:"tgl_akhir"`
	CreateAt         string                   `json:"create_at"`
	StatusTTD        string                   `json:"status_ttd"`
	UserNIP          string                   `json:"user_nip"`
	UserName         string                   `json:"user_name"`
	UserNoTelp       string                   `json:"user_no_telp"`
	UserEmail        string                   `json:"user_email"`
	Participans      []ParticipanJoinUserFoto `json:"participans"`
}

// Participan
type ParticipanJoinUser struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	NIP          string `json:"nip"`
	Name         string `json:"name"`
	NoTelp       string `json:"no_telp"`
	Email        string `json:"email"`
}

type ParticipanJoinUserFoto struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	NIP          string `json:"nip"`
	Name         string `json:"name"`
	NoTelp       string `json:"no_telp"`
	Email        string `json:"email"`
	NameGambar   string `json:"name_gambar"`
	Gambar       string `json:"gambar"`
	Lokasi       string `json:"lokasi"`
	Koordinat    string `json:"koordinat"`
}

// Laporan
type LaporanJoinApproved struct {
	Id           int    `json:"id"`
	SuratTugasId int    `json:"surat_tugas_id"`
	UserId       int    `json:"user_id"`
	DokName      string `json:"dok_name"`
	DokPDF       string `json:"dok_pdf"`
	CreateAt     string `json:"create_at"`
	Status       string `json:"status"`
	Message      string `json:"message"`
}

type LaporanAkAngg struct {
	SuratTugasId int    `json:"surat_tugas_id"`
	UserId       int    `json:"user_id"`
	DokAkName    string `json:"dok_ak_name"`
	DokAkPDF     string `json:"dok_ak_pdf"`
	DokAgName    string `json:"dok_ag_name"`
	DokAgPDF     string `json:"dok_ag_pdf"`
}
