package surattugasreqres

type SuratTugasResponse struct {
	Id               int    `json:"id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     int    `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
}

type SuratTugasJOINApprovedResponse struct {
	Id               int    `json:"id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     int    `json:"jenis_program"`
	DokumenName      string `json:"dokumen_name"`
	DokumenPDF       string `json:"dokumen_pdf"`
	DokPendukungName string `json:"dok_pendukung_name"`
	DokPendukungPdf  string `json:"dok_pendukung_pdf"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
	Status           int    `json:"status"`
}
