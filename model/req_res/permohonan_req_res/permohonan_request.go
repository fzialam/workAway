package permohonanreqres

type PermohonanRequest struct {
	UserPemohonId    int    `validate:"required" json:"user_pemohon_id"`
	LokasiTujuan     string `validate:"required" json:"lokasi_tujuan"`
	JenisProgram     string `validate:"required" json:"jenis_program"`
	DokPendukungName string `validate:"required" json:"dok_pendukung_name"`
	DokPendukungPdf  string `validate:"required" json:"dok_pendukung_pdf"`
	ParticipansId    []int  `validate:"" json:"participans_id"`
	TglAwal          string `validate:"required" json:"tgl_awal"`
	TglAkhir         string `validate:"required" json:"tgl_akhir"`
}
