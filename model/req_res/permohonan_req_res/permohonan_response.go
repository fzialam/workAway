package permohonanreqres

type PermohonanResponse struct {
	UserPemohonId    int    `json:"user_pemohon_id"`
	LokasiTujuan     string `json:"lokasi_tujuan"`
	JenisProgram     string `json:"jenis_program"`
	DokPendukungName string `json:"dok_pendukung_name"`
	ParticipansId    []int  `json:"participans_id"`
	TglAwal          string `json:"tgl_awal"`
	TglAkhir         string `json:"tgl_akhir"`
}
