package penugasanreqres

type PenugasanRequest struct {
	Tipe          int    `validate:"required" json:"tipe"`
	UserKetuaId   int    `validate:"required" json:"user_ketua_id"`
	LokasiTujuan  string `validate:"required" json:"lokasi_tujuan"`
	JenisProgram  string `validate:"required" json:"jenis_program"`
	ParticipansId []int  `validate:"" json:"participans_id"`
	TglAwal       string `validate:"required" json:"tgl_awal"`
	TglAkhir      string `validate:"required" json:"tgl_akhir"`
}

type ApprovedRequest struct {
	SuratTugasId int    `validate:"required" json:"surat_tugas_id"`
	Status       string `validate:"required" json:"status"`
	StatusTTD    string `validate:"required" json:"status_ttd"`
}
