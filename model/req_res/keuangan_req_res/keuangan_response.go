package keuanganreqres

type UploadRincianAnggaranResponse struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	DokName      string `validate:"required" json:"dok_name"`
	DokPDF       string `validate:"required" json:"dok_pdf"`
}

type SuratTugasJOINApprovedUserOtherIdResponse struct {
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

type SPPDRincian struct {
	IsCreated int `json:"iscreated"`
	Approved  int `json:"approved"`
	Reject    int `json:"reject"`
}

type Laporan struct {
	Belum    int `json:"belum"`
	Approved int `json:"approved"`
	Reject   int `json:"reject"`
}

type IndexKeuangan struct {
	SPPDRincian SPPDRincian `json:"sppd_rincian"`
	Laporan     Laporan     `json:"laporan"`
}
